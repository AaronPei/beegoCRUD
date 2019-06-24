package controllers

import (
	"encoding/json"
	"fmt"
	"learn/beegoTest/nodeApi/models"
	"time"

	"github.com/astaxie/beego"
)

type NodesController struct {
	beego.Controller
}

// @router / [post]
func (n *NodesController) Post() {
	var addNodeStruct models.AddNodeStruct
	json.Unmarshal(n.Ctx.Input.RequestBody, &addNodeStruct)
	if addNodeStruct.Name == "" ||
		addNodeStruct.Idc == "" ||
		addNodeStruct.Role == "" ||
		addNodeStruct.Branch == "" {
		panic("Name,Idc,Role,Branch are required!")
	}
	uid := models.AddNode(addNodeStruct)
	n.Data["json"] = map[string]int64{"id": uid}
	n.ServeJSON()
}

// @router / [get]
func (n *NodesController) GetAll() {
	nodes := models.GetAllNodes()
	n.Data["json"] = nodes
	n.ServeJSON()
}

// @Title Get
// @Description get node by id
// @Param	id		path 	string	true		'The key for node'
// @Success 200 {object} models.Node
// @Failure 403 :id is empty
// @router /:id [get]
func (n *NodesController) Get() {
	id := n.GetString(":id")
	fmt.Println("Id = " + id)
	if id != "" {
		node, err := models.GetNode(id)
		if err != nil {
			n.Data["json"] = err.Error()
		} else {
			n.Data["json"] = node
		}
	}
	n.ServeJSON()
}

// @Title Update
// @Description update node by id
// @Param	id		path 	string	true		'The id you want to udpate'
// @Param	body	path 	models.AddNodeStruct	true		'The body for Node'
// @Success 200 {object} models.AddNodeStruct
// @Failure 403 :id is not int
// @router /:id [put]
func (n *NodesController) Update() {
	id := n.GetString(":id")
	if id != "" {
		var addNodeStruct models.AddNodeStruct
		json.Unmarshal(n.Ctx.Input.RequestBody, &addNodeStruct)
		if addNodeStruct.Name == "" ||
			addNodeStruct.Idc == "" ||
			addNodeStruct.Role == "" ||
			addNodeStruct.Branch == "" {
			panic("Name,Idc,Role,Branch are required!")
		}
		layout := "2006-01-02 15:04:05"
		str := addNodeStruct.Submission_at
		t, err := time.Parse(layout, str)
		if err != nil {
			panic(err)
		}
		addNodeStruct.Submission_date = t
		isSuc, err := models.UpdateNode(id, &addNodeStruct)
		if err != nil {
			n.Data["json"] = err.Error()
		} else {
			n.Data["json"] = map[string]bool{"update": isSuc}
		}
	}
	n.ServeJSON()
}

// @Title Delete
// @Description delete node by id
// @Param	id		path 	string	true		'The key for node'
// @Success 200 delete success
// @Failure 403 :id is empty
// @router /:id [delete]
func (n *NodesController) Delete() {
	id := n.GetString(":id")
	var err error
	err = models.DeleteNode(id)
	if err == nil {
		n.Data["json"] = "delete success!"
	} else {
		n.Data["json"] = "delete failed!"
	}
	n.ServeJSON()
}
