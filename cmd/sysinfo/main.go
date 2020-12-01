package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	"github.com/lonelypale/gotools/network"
)

func main() {
	collet()
}

func collet() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")

	if len(c) > 1 {
		for _, subCpu := range c {
			modelname := subCpu.ModelName
			cores := subCpu.Cores
			fmt.Printf("CPU       : %v   %v cores \n", modelname, cores)
		}
	} else {
		subCpu := c[0]
		modelname := subCpu.ModelName
		cores := subCpu.Cores
		fmt.Printf("CPU       : %v   %v cores \n", modelname, cores)
	}
	fmt.Printf("CPU Used  : used %f%% \n", cc[0])
	fmt.Printf("Mem       : %vMB  Free: %vMB  Used: %v Usage: %f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	fmt.Printf("Disk      : %vGB  Free: %vGB  Usage: %f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("Network   : %vbytes / %vbytes\n", nv[0].BytesRecv, nv[0].BytesSent)
	fmt.Printf("SystemBoot: %v\n", btime)
	fmt.Printf("OS        : %v(%v)   %v  \n", n.Platform, n.PlatformFamily, n.PlatformVersion)
	fmt.Printf("Hostname  : %v  \n", n.Hostname)

	ip, err := network.GetLocalIP()
	if err != nil {
		panic(err)
	}
	fmt.Printf("IP        : %v  \n", ip)
}
