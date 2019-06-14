// Copyright (c) 2017, NVIDIA CORPORATION. All rights reserved.

package main

import (
	"github.com/jaypipes/ghw"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
	"log"
)

func check(err error) {
	if err != nil {
		log.Panicln("Fatal:", err)
	}
}

func getDevices() []*pluginapi.Device {
	pci, err := ghw.PCI()
	check(err)

	var devs []*pluginapi.Device
	for _, pciInfo := range pci.ListDevices() {
		// Only support Nvidia and AMD GPU
		if (pciInfo.Vendor.ID == "10de" || pciInfo.Vendor.ID == "1002") && pciInfo.Class.ID == "03" {
			log.Println("Found GPU:", pciInfo)
			devs = append(devs, &pluginapi.Device{
				ID:     pciInfo.Address,
				Health: pluginapi.Healthy,
			})
		}
	}

	return devs
}

func deviceExists(devs []*pluginapi.Device, id string) bool {
	for _, d := range devs {
		if d.ID == id {
			return true
		}
	}
	return false
}

//func watchXIDs(ctx context.Context, devs []*pluginapi.Device, xids chan<- *pluginapi.Device) {
//	eventSet := nvml.NewEventSet()
//	defer nvml.DeleteEventSet(eventSet)
//
//	for _, d := range devs {
//		err := nvml.RegisterEventForDevice(eventSet, nvml.XidCriticalError, d.ID)
//		if err != nil && strings.HasSuffix(err.Error(), "Not Supported") {
//			log.Printf("Warning: %s is too old to support healthchecking: %s. Marking it unhealthy.", d.ID, err)
//
//			xids <- d
//			continue
//		}
//
//		if err != nil {
//			log.Panicln("Fatal:", err)
//		}
//	}
//
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		default:
//		}
//
//		e, err := nvml.WaitForEvent(eventSet, 5000)
//		if err != nil && e.Etype != nvml.XidCriticalError {
//			continue
//		}
//
//		// FIXME: formalize the full list and document it.
//		// http://docs.nvidia.com/deploy/xid-errors/index.html#topic_4
//		// Application errors: the GPU should still be healthy
//		if e.Edata == 31 || e.Edata == 43 || e.Edata == 45 {
//			continue
//		}
//
//		if e.UUID == nil || len(*e.UUID) == 0 {
//			// All devices are unhealthy
//			for _, d := range devs {
//				xids <- d
//			}
//			continue
//		}
//
//		for _, d := range devs {
//			if d.ID == *e.UUID {
//				xids <- d
//			}
//		}
//	}
//}
