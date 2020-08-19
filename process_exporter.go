package main

import (
	//"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/process"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ProcessCollector struct {
	proCpuUsage *prometheus.Desc // cpu使用率
	proMemUsage *prometheus.Desc //mem使用率
	proNumFds *prometheus.Desc  //文件打开数
	proReadBytes *prometheus.Desc //读字节数
	proWriteBytes *prometheus.Desc //写字节数
}



func NewProcessCollector() prometheus.Collector {

	return &ProcessCollector{
		proCpuUsage:prometheus.NewDesc(

			"process_cpu_percent",
			"Get the cpu usage rate of a single process",
			[]string{"ProName","ProPid","ProCreateTime","ProUserName","proCwd"},nil),
		proMemUsage:prometheus.NewDesc(prometheus.BuildFQName(
			"process",
			"mem",
			"percent"),
			"Get the mem usage rate of a single process",
			[]string{"ProName","ProPid","ProCreateTime","ProUserName","proCwd"},nil),
		proNumFds:prometheus.NewDesc(prometheus.BuildFQName(
			"process",
			"open_files",
			"num"),
			"Number of open files per process",
			[]string{"ProName","ProPid","ProCreateTime","ProUserName","proCwd"},nil),
		proReadBytes:prometheus.NewDesc(prometheus.BuildFQName("process",
			"read",
			"bytes"),
			"Number of bytes read by the process",
			[]string{"ProName","ProPid","ProCreateTime","ProUserName","proCwd"},nil),
		proWriteBytes:prometheus.NewDesc(prometheus.BuildFQName("process",
			"write",
			"bytes"),
			"Number of bytes write by the process",
			[]string{"ProName","ProPid","ProCreateTime","ProUserName","proCwd"},nil),
	}
}

//实现采集器Describe接口
func (n *ProcessCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.proCpuUsage
	ch <- n.proMemUsage
	ch <- n.proNumFds
	ch <- n.proReadBytes
	ch <- n.proWriteBytes
}

func NewPidsWithContext(path string) ([]int32, error){
	var ret []int32

	d, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	fnames, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	for _, fname := range fnames {
		pid, err := strconv.ParseInt(fname, 10, 32)
		if err != nil {
			// if not numeric name, just skip
			continue
		}
		//statPath := common.HostProc(strconv.Itoa(int(pid)), "status")
		statPath := path + strconv.Itoa(int(pid)) + "/" + "status"
		contents, err2 := ioutil.ReadFile(statPath)
		if err2 != nil {
			fmt.Println("No status file . error: ",err2)
			continue
		}
		lines := strings.Split(string(contents), "\n")
		for i:=1;i <11;i++{
			text := lines[i]
			tabParts := strings.SplitN(text, "\t", 2)
			value := tabParts[1]
			switch strings.TrimRight(tabParts[0], ":"){
			case "PPid", "Ppid":
				pval, _ := strconv.ParseInt(value, 10, 32)
				//p.parent = int32(pval)
				if pval != int64(2){
					ret = append(ret, int32(pid))
				}
			}
		}
	}
	//pids,_ :=process.Pids()
	return ret, nil
}

// Collect returns the current state of all metrics of the collector.
//实现采集器Collect接口,真正采集动作
func (n *ProcessCollector) Collect(ch chan<- prometheus.Metric) {
	pids,_ := NewPidsWithContext("/proc/")
	for _,pid := range pids{
		p, _ := process.NewProcess(int32(pid))
		tempnum,_  := p.Tgid()
		if tempnum != int32(0){
			useCpuUsage, _ := p.CPUPercent() // cpu使用率
			useMemUsage, _ := p.MemoryPercent()
			useCreateTime, _ := p.CreateTime()
			//usePid, _ := p.Tgid()
			usePid := int32(pid)
			useName, _ := p.Name()
			useUserName, _ := p.Username()
			useNumFds, _ := p.NumFDs()
			usePwd,_ := p.Cwd()

			iostatus,ioerr := p.IOCounters()
			if ioerr != nil{
				fmt.Println("No io information for this process. error: ",ioerr)
				continue
			}
			useReadBytes := iostatus.ReadBytes
			useWriteBytes := iostatus.WriteBytes

			useNumFdsString := strconv.FormatInt(int64(useNumFds), 10)
			useNumFdsFloat64, _ := strconv.ParseFloat(useNumFdsString, 64)
			ch <- prometheus.MustNewConstMetric(n.proCpuUsage, prometheus.GaugeValue, useCpuUsage, string(useName), strconv.Itoa(int(usePid)), strconv.FormatInt(useCreateTime, 10), string(useUserName),usePwd)
			ch <- prometheus.MustNewConstMetric(n.proMemUsage, prometheus.GaugeValue, float64(useMemUsage), string(useName), strconv.Itoa(int(usePid)), strconv.FormatInt(useCreateTime, 10), string(useUserName),usePwd)
			ch <- prometheus.MustNewConstMetric(n.proNumFds, prometheus.GaugeValue, useNumFdsFloat64, string(useName), strconv.Itoa(int(usePid)), strconv.FormatInt(useCreateTime, 10), string(useUserName),usePwd)
			ch <- prometheus.MustNewConstMetric(n.proReadBytes,prometheus.GaugeValue,float64(useReadBytes),string(useName), strconv.Itoa(int(usePid)), strconv.FormatInt(useCreateTime, 10), string(useUserName),usePwd)
			ch <- prometheus.MustNewConstMetric(n.proWriteBytes,prometheus.GaugeValue,float64(useWriteBytes),string(useName), strconv.Itoa(int(usePid)), strconv.FormatInt(useCreateTime, 10), string(useUserName),usePwd)
			//fmt.Println(pid,usePid,useName,useCpuUsage,useMemUsage,useNumFds,useUserName,useCreateTime)
			//fmt.Println(useReadBytes,useWriteBytes)
		}
		}

	}




func init()  {
	//注册自身采集器
	prometheus.MustRegister(NewProcessCollector())
}


func main() {
	var (
		listenAddress = kingpin.Flag(
			"web.listen-address",
			"Address on which to expose metrics and web interface.",
		).Default(":19124").String()
		metricsPath = kingpin.Flag(
			"web.telemetry-path",
			"Path under which to expose metrics.",
		).Default("/metrics").String()

	)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	http.Handle(*metricsPath, promhttp.Handler())
	//fmt.Println("开始启动")
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		fmt.Printf("frror occur when start server %v", err)
	}

}
