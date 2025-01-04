// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package miner_core

import (
	"context"
	"fmt"
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
	"github.com/qcq1/rpc_miner_core/kitex_gen/base"
	"strings"
)

type Item struct {
	Id    int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Title string `thrift:"title,2" frugal:"2,default,string" json:"title"`
	Stock int64  `thrift:"stock,3" frugal:"3,default,i64" json:"stock"`
}

func NewItem() *Item {
	return &Item{}
}

func (p *Item) InitDefault() {
}

func (p *Item) GetId() (v int64) {
	return p.Id
}

func (p *Item) GetTitle() (v string) {
	return p.Title
}

func (p *Item) GetStock() (v int64) {
	return p.Stock
}
func (p *Item) SetId(val int64) {
	p.Id = val
}
func (p *Item) SetTitle(val string) {
	p.Title = val
}
func (p *Item) SetStock(val int64) {
	p.Stock = val
}

var fieldIDToName_Item = map[int16]string{
	1: "id",
	2: "title",
	3: "stock",
}

func (p *Item) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Item[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Item) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *Item) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Title = _field
	return nil
}
func (p *Item) ReadField3(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Stock = _field
	return nil
}

func (p *Item) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("Item"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Item) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Item) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("title", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Item) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("stock", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Stock); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Item) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Item(%+v)", *p)

}

func (p *Item) DeepEqual(ano *Item) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Title) {
		return false
	}
	if !p.Field3DeepEqual(ano.Stock) {
		return false
	}
	return true
}

func (p *Item) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *Item) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Title, src) != 0 {
		return false
	}
	return true
}
func (p *Item) Field3DeepEqual(src int64) bool {

	if p.Stock != src {
		return false
	}
	return true
}

type GetItemReq struct {
	Id   int64      `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
	Base *base.Base `thrift:"base,255" frugal:"255,default,base.Base" json:"base"`
}

func NewGetItemReq() *GetItemReq {
	return &GetItemReq{}
}

func (p *GetItemReq) InitDefault() {
}

func (p *GetItemReq) GetId() (v int64) {
	return p.Id
}

var GetItemReq_Base_DEFAULT *base.Base

func (p *GetItemReq) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return GetItemReq_Base_DEFAULT
	}
	return p.Base
}
func (p *GetItemReq) SetId(val int64) {
	p.Id = val
}
func (p *GetItemReq) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_GetItemReq = map[int16]string{
	1:   "id",
	255: "base",
}

func (p *GetItemReq) IsSetBase() bool {
	return p.Base != nil
}

func (p *GetItemReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetItemReq[fieldId]))
}

func (p *GetItemReq) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *GetItemReq) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBase()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}

func (p *GetItemReq) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemReq) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Base.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemReq(%+v)", *p)

}

func (p *GetItemReq) DeepEqual(ano *GetItemReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *GetItemReq) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *GetItemReq) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type GetItemResp struct {
	Item     *Item          `thrift:"item,1" frugal:"1,default,Item" json:"item"`
	BaseResp *base.BaseResp `thrift:"baseResp,255" frugal:"255,default,base.BaseResp" json:"baseResp"`
}

func NewGetItemResp() *GetItemResp {
	return &GetItemResp{}
}

func (p *GetItemResp) InitDefault() {
}

var GetItemResp_Item_DEFAULT *Item

func (p *GetItemResp) GetItem() (v *Item) {
	if !p.IsSetItem() {
		return GetItemResp_Item_DEFAULT
	}
	return p.Item
}

var GetItemResp_BaseResp_DEFAULT *base.BaseResp

func (p *GetItemResp) GetBaseResp() (v *base.BaseResp) {
	if !p.IsSetBaseResp() {
		return GetItemResp_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *GetItemResp) SetItem(val *Item) {
	p.Item = val
}
func (p *GetItemResp) SetBaseResp(val *base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_GetItemResp = map[int16]string{
	1:   "item",
	255: "baseResp",
}

func (p *GetItemResp) IsSetItem() bool {
	return p.Item != nil
}

func (p *GetItemResp) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *GetItemResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetItemResp) ReadField1(iprot thrift.TProtocol) error {
	_field := NewItem()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Item = _field
	return nil
}
func (p *GetItemResp) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *GetItemResp) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("item", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Item.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemResp) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("baseResp", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemResp(%+v)", *p)

}

func (p *GetItemResp) DeepEqual(ano *GetItemResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Item) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *GetItemResp) Field1DeepEqual(src *Item) bool {

	if !p.Item.DeepEqual(src) {
		return false
	}
	return true
}
func (p *GetItemResp) Field255DeepEqual(src *base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}

type ItemService interface {
	GetItem(ctx context.Context, req *GetItemReq) (r *GetItemResp, err error)
}

type ItemServiceGetItemArgs struct {
	Req *GetItemReq `thrift:"req,1" frugal:"1,default,GetItemReq" json:"req"`
}

func NewItemServiceGetItemArgs() *ItemServiceGetItemArgs {
	return &ItemServiceGetItemArgs{}
}

func (p *ItemServiceGetItemArgs) InitDefault() {
}

var ItemServiceGetItemArgs_Req_DEFAULT *GetItemReq

func (p *ItemServiceGetItemArgs) GetReq() (v *GetItemReq) {
	if !p.IsSetReq() {
		return ItemServiceGetItemArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *ItemServiceGetItemArgs) SetReq(val *GetItemReq) {
	p.Req = val
}

var fieldIDToName_ItemServiceGetItemArgs = map[int16]string{
	1: "req",
}

func (p *ItemServiceGetItemArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ItemServiceGetItemArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ItemServiceGetItemArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGetItemReq()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *ItemServiceGetItemArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItem_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ItemServiceGetItemArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemArgs(%+v)", *p)

}

func (p *ItemServiceGetItemArgs) DeepEqual(ano *ItemServiceGetItemArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *ItemServiceGetItemArgs) Field1DeepEqual(src *GetItemReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type ItemServiceGetItemResult struct {
	Success *GetItemResp `thrift:"success,0,optional" frugal:"0,optional,GetItemResp" json:"success,omitempty"`
}

func NewItemServiceGetItemResult() *ItemServiceGetItemResult {
	return &ItemServiceGetItemResult{}
}

func (p *ItemServiceGetItemResult) InitDefault() {
}

var ItemServiceGetItemResult_Success_DEFAULT *GetItemResp

func (p *ItemServiceGetItemResult) GetSuccess() (v *GetItemResp) {
	if !p.IsSetSuccess() {
		return ItemServiceGetItemResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ItemServiceGetItemResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetItemResp)
}

var fieldIDToName_ItemServiceGetItemResult = map[int16]string{
	0: "success",
}

func (p *ItemServiceGetItemResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ItemServiceGetItemResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ItemServiceGetItemResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ItemServiceGetItemResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewGetItemResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *ItemServiceGetItemResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItem_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ItemServiceGetItemResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *ItemServiceGetItemResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemResult(%+v)", *p)

}

func (p *ItemServiceGetItemResult) DeepEqual(ano *ItemServiceGetItemResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *ItemServiceGetItemResult) Field0DeepEqual(src *GetItemResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
