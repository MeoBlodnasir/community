/*
 * Nanocloud Enterprise, a comprehensive platform to turn any application
 * into a cloud solution.
 *
 * Copyright (C) 2016 Nanocloud Software
 *
 * This file is part of Nanocloud Enterprise
 */

package azure

import (
	"bitbucket.org/nanocloud/enterprise/nanocloud/models/azure"
	log "github.com/Sirupsen/logrus"

	"github.com/Nanocloud/community/nanocloud/iaas"
	"github.com/Nanocloud/community/nanocloud/router"
)

type hash map[string]interface{}

func ListRunningVM(req router.Request) (*router.Response, error) {
	response, err := iaas.ListVM()
	if err != nil {
		log.Error("Unable to retrieve VM states list")
		return router.JSONResponse(500, hash{
			"error": "Unable te retrieve states of VMs: " + err.Error(),
		}), err
	}
	vmList := iaas.CheckVMStates(response)
	return router.JSONResponse(200, vmList), nil
}

func DownloadVM(req router.Request) (*router.Response, error) {
	return router.JSONResponse(400, hash{
		"error": "Unimplemented",
	}), nil
}

func StartVM(req router.Request) (*router.Response, error) {
	var params = map[string]string{
		"name": req.Params["id"],
	}

	if params["name"] == "" {
		return router.JSONResponse(400, hash{
			"error": "No VM name provided",
		}), nil
	}

	vm, err := iaas.GetVM(params["name"])

	err := vm.Start()
	if err != nil {
		log.Error("Error while starting VM")
		return router.JSONResponse(500, hash{
			"error": "Unable to start the specified VM",
		}), err
	}

	return router.JSONResponse(200, hash{
		"success": true,
	}), nil
}

func StopVM(req router.Request) (*router.Response, error) {
	var params = map[string]string{
		"Name": req.Params["id"],
	}

	if params["Name"] == "" {
		return router.JSONResponse(400, hash{
			"error": "No VM name provided",
		}), nil
	}

	err := azure.Stop(params["Name"])
	if err != nil {
		return router.JSONResponse(500, hash{
			"error": "Unable to stop the specified VM",
		}), err
	}

	return router.JSONResponse(200, hash{
		"success": true,
	}), nil
}
