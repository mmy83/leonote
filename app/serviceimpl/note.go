/*
@Time : 2020/4/25 1:01 下午
@Author : mmy83
@File : note.go
@Software: GoLand
*/

package serviceimpl

import (
	"leonote/app/model"
	"leonote/app/service"
	"leonote/config"
	"leonote/database"
	"log"
	"strconv"
	"strings"
)

func NewNote() service.Note {
	return &note{
		pageSize: config.CfgPage.GetInt("pageSize"),
	}
}

type note struct {
	pageSize int
}

func (n *note) GetList(userId int64,notebookId int64) ([]*model.Note,error){
	notes := make([]*model.Note,0)
	err := database.Engine.Where("notebook_id=? and user_id=?",notebookId,userId).Limit(n.pageSize,0).Find(&notes)
	return notes,err
}

func (n *note) GetNote(id int64, userId int64)(*model.Note,bool,error){

	note := &model.Note{
		Id:     id,
		UserId: userId,
	}
	has,err := database.Engine.Get(note)
	return note,has, err
}



func (n *note) CreateNote(note *model.Note) (int64,error) {
	session := database.Engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return 0, err
	}

	affected, err := session.Insert(note)

	log.Printf("note into : %v\n",note)
	if err != nil {
		session.Rollback()
		return affected, err
	}

	//tags
	tagNames := strings.Split(note.Tags,"|")
	tagNameMap := make(map[string]string)
	for _,v:= range tagNames {
		tagNameMap[v] = v
	}
	tagNames = tagNames[:0]

	for _, v := range tagNameMap{
		tagNames = append(tagNames, v)
	}

	if len(tagNames) <= 0 || (len(tagNames)==1 && (tagNames[0]=="" || tagNames[0]=="")){

		session.Commit()
		return affected, err
	}
	tagNameStr := strings.Join(tagNames,`","`)

	updateTags := make([]*model.Tag,0)
	err = session.Where("user_id = ? and tag_name in (\""+tagNameStr+"\")",note.UserId).Find(&updateTags)
	if err != nil {
		session.Rollback()
		return 0,err
	}

	log.Printf("select updateTags: %s\n",updateTags)

	updateTagMap := make(map[string]string)
	updateTagIds := make([]string,0)
	for _,v := range updateTags {
		updateTagMap[v.TagName] = v.TagName
		updateTagIds = append(updateTagIds, strconv.FormatInt(v.Id,10))
	}

	log.Printf("+1: %s\n",updateTagIds)
	//update
	if len(updateTagIds)>0 {
		tag := new(model.Tag)
		sql := "update "+tag.TableName()+" set count=count+1 where id in ("+
			strings.Join(updateTagIds,",") +")"
		log.Printf("sql: %s\n",sql)
		_,err = session.Exec(sql)
		if err != nil {
			session.Rollback()
			return 0,err
		}
	}

	//insert
	insertTags := make([]*model.Tag,0)

	insertTagNames := make([]string,0)
	for _,val := range tagNames{
		if _,ok := updateTagMap[val];ok {
			continue
		}
		insertTags = append(insertTags, &model.Tag{
			TagName: val,
			Count: 1,
			UserId: note.UserId,
		})
		insertTagNames = append(insertTagNames,val)
	}

	_,err = session.Insert(&insertTags)
	if err != nil {
		session.Rollback()
		return affected, err
	}
	insertTags = insertTags[:0]
	insertTagNamesStr := strings.Join(insertTagNames,`","`)
	session.Where("user_id= ? and tag_name in ( \""+insertTagNamesStr+"\" )",note.UserId).Find(&insertTags)

	log.Printf("insert tags: %s\n",insertTags)

	noteTags := make([]*model.NoteTag,0)
	for _,value := range updateTags {
		noteTags = append(noteTags, &model.NoteTag{
			NoteId: note.Id,
			TagId: value.Id,
			UserId: note.UserId,
		})
	}
	for _,value := range insertTags {
		noteTags = append(noteTags, &model.NoteTag{
			NoteId: note.Id,
			TagId: value.Id,
			UserId: note.UserId,
		})
	}

	_,err = session.Insert(noteTags)
	if err != nil {
		session.Rollback()
		return affected, err
	}

	log.Printf("insert notetags: %s\n",noteTags)

	session.Commit()

	return affected, err



}