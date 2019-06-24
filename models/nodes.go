package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Node struct {
	Id              int       `orm:"pk;auto"`
	Name            string    `orm:"size(100)"`
	Role            string    `orm:"size(100)"`
	Idc             string    `orm:"size(100)"`
	Branch          string    `orm:"size(100)"`
	Submission_date time.Time `orm:"type(datetime)"`
}

type AddNodeStruct struct {
	Name            string `json:"Name"`
	Role            string
	Idc             string
	Branch          string
	Submission_at   string
	Submission_date time.Time
}

func init() {
	orm.RegisterModel(new(Node))
}

func AddNode(addNodeStruct AddNodeStruct) int64 {
	o := orm.NewOrm()
	var node Node
	node.Name = addNodeStruct.Name
	node.Role = addNodeStruct.Role
	node.Idc = addNodeStruct.Idc
	node.Branch = addNodeStruct.Branch
	node.Submission_date = time.Now()
	id, err := o.Insert(&node)
	if err != nil {
		panic(err)
	}
	return id
}

func GetAllNodes() []Node {
	var nodes []Node
	var o = orm.NewOrm()
	var maps []orm.Params
	num, err := o.QueryTable("node").Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			var each = Node{}
			idInt64 := m["Id"].(int64)
			strIdInt64 := strconv.FormatInt(idInt64, 10)
			idInt16, _ := strconv.Atoi(strIdInt64)
			each.Id = idInt16
			each.Name = m["Name"].(string)
			each.Role = m["Role"].(string)
			each.Idc = m["Idc"].(string)
			each.Branch = m["Branch"].(string)
			each.Submission_date = m["Submission_date"].(time.Time)
			nodes = append(nodes, each)
		}
	}
	return nodes
}

func GetNode(id string) (Node, error) {
	o := orm.NewOrm()
	var err error
	var node Node
	err = o.QueryTable("node").Filter("id", id).One(&node)
	if err == orm.ErrMultiRows {
		// Have multiple records
		fmt.Printf("Returned Multi Rows Not One")
		return node, errors.New("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// No result
		fmt.Printf("Not row found")
		return node, errors.New("Not row found")
	}
	return node, nil
}

func UpdateNode(id string, addNodeStruct *AddNodeStruct) (bool, error) {
	o := orm.NewOrm()
	exist := o.QueryTable("node").Filter("id", id).Exist()
	if exist {
		num, err := o.QueryTable("node").Filter("id", id).Update(orm.Params{
			"Name":            addNodeStruct.Name,
			"Role":            addNodeStruct.Role,
			"Idc":             addNodeStruct.Idc,
			"Branch":          addNodeStruct.Branch,
			"Submission_date": addNodeStruct.Submission_date,
		})
		if err != nil {
			return false, err
		}
		fmt.Println("num:%s\n", num)
		return true, nil
	} else {
		return false, errors.New("Data not exist")
	}
}

func DeleteNode(id string) error {
	o := orm.NewOrm()
	exist := o.QueryTable("node").Filter("id", id).Exist()
	if exist {
		num, err := o.QueryTable("node").Filter("id", id).Delete()
		if err != nil {
			fmt.Println("Delete DATA failed")
			return err
		}
		if num < 1 {
			return errors.New("Delete Failed!DATA is not exist")
		} else {
			fmt.Println("Delete success")
			return nil
		}
	} else {
		return errors.New("Data not exist")
	}
}
