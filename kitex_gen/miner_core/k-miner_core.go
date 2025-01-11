// Code generated by Kitex v0.12.1. DO NOT EDIT.

package miner_core

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"

	"github.com/qcq1/rpc_miner_core/kitex_gen/common"
)

var (
	_ = common.KitexUnusedProtection
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *Job) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Job[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *Job) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Id = _field
	return offset, nil
}

func (p *Job) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Name = _field
	return offset, nil
}

func (p *Job) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Description = _field
	return offset, nil
}

func (p *Job) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.CreatedBy = _field
	return offset, nil
}

func (p *Job) FastReadField5(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.UpdatedBy = _field
	return offset, nil
}

func (p *Job) FastReadField6(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.CreatedAt = _field
	return offset, nil
}

func (p *Job) FastReadField7(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.UpdatedAt = _field
	return offset, nil
}

func (p *Job) FastReadField8(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Extra = _field
	return offset, nil
}

func (p *Job) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *Job) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
		offset += p.fastWriteField7(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField8(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *Job) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *Job) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.Id)
	return offset
}

func (p *Job) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Name)
	return offset
}

func (p *Job) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 3)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Description)
	return offset
}

func (p *Job) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 4)
	offset += thrift.Binary.WriteI64(buf[offset:], p.CreatedBy)
	return offset
}

func (p *Job) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 5)
	offset += thrift.Binary.WriteI64(buf[offset:], p.UpdatedBy)
	return offset
}

func (p *Job) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 6)
	offset += thrift.Binary.WriteI64(buf[offset:], p.CreatedAt)
	return offset
}

func (p *Job) fastWriteField7(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 7)
	offset += thrift.Binary.WriteI64(buf[offset:], p.UpdatedAt)
	return offset
}

func (p *Job) fastWriteField8(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 8)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Extra)
	return offset
}

func (p *Job) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *Job) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Name)
	return l
}

func (p *Job) field3Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Description)
	return l
}

func (p *Job) field4Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *Job) field5Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *Job) field6Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *Job) field7Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *Job) field8Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Extra)
	return l
}

func (p *QueryJobListReq) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPageNum bool = false
	var issetPageSize bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPageNum = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPageSize = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetPageNum {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetPageSize {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_QueryJobListReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_QueryJobListReq[fieldId]))
}

func (p *QueryJobListReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.PageNum = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.PageSize = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field *JobColumn
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		tmp := JobColumn(v)
		_field = &tmp
	}
	p.OrderBy = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field *common.Order
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		tmp := common.Order(v)
		_field = &tmp
	}
	p.Order = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField5(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Id = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField6(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.CreatedBy = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField7(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.CreatedAtStart = _field
	return offset, nil
}

func (p *QueryJobListReq) FastReadField8(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.CreatedAtEnd = _field
	return offset, nil
}

func (p *QueryJobListReq) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *QueryJobListReq) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
		offset += p.fastWriteField7(buf[offset:], w)
		offset += p.fastWriteField8(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *QueryJobListReq) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *QueryJobListReq) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.PageNum)
	return offset
}

func (p *QueryJobListReq) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 2)
	offset += thrift.Binary.WriteI64(buf[offset:], p.PageSize)
	return offset
}

func (p *QueryJobListReq) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetOrderBy() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 3)
		offset += thrift.Binary.WriteI32(buf[offset:], int32(*p.OrderBy))
	}
	return offset
}

func (p *QueryJobListReq) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetOrder() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 4)
		offset += thrift.Binary.WriteI32(buf[offset:], int32(*p.Order))
	}
	return offset
}

func (p *QueryJobListReq) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetId() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 5)
		offset += thrift.Binary.WriteI64(buf[offset:], *p.Id)
	}
	return offset
}

func (p *QueryJobListReq) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetCreatedBy() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 6)
		offset += thrift.Binary.WriteI64(buf[offset:], *p.CreatedBy)
	}
	return offset
}

func (p *QueryJobListReq) fastWriteField7(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetCreatedAtStart() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 7)
		offset += thrift.Binary.WriteI64(buf[offset:], *p.CreatedAtStart)
	}
	return offset
}

func (p *QueryJobListReq) fastWriteField8(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetCreatedAtEnd() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 8)
		offset += thrift.Binary.WriteI64(buf[offset:], *p.CreatedAtEnd)
	}
	return offset
}

func (p *QueryJobListReq) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *QueryJobListReq) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *QueryJobListReq) field3Length() int {
	l := 0
	if p.IsSetOrderBy() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I32Length()
	}
	return l
}

func (p *QueryJobListReq) field4Length() int {
	l := 0
	if p.IsSetOrder() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I32Length()
	}
	return l
}

func (p *QueryJobListReq) field5Length() int {
	l := 0
	if p.IsSetId() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I64Length()
	}
	return l
}

func (p *QueryJobListReq) field6Length() int {
	l := 0
	if p.IsSetCreatedBy() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I64Length()
	}
	return l
}

func (p *QueryJobListReq) field7Length() int {
	l := 0
	if p.IsSetCreatedAtStart() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I64Length()
	}
	return l
}

func (p *QueryJobListReq) field8Length() int {
	l := 0
	if p.IsSetCreatedAtEnd() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I64Length()
	}
	return l
}

func (p *QueryJobListResp) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_QueryJobListResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *QueryJobListResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := thrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	_field := make([]*Job, 0, size)
	values := make([]Job, size)
	for i := 0; i < size; i++ {
		_elem := &values[i]
		_elem.InitDefault()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		_field = append(_field, _elem)
	}
	p.JobList = _field
	return offset, nil
}

func (p *QueryJobListResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Total = _field
	return offset, nil
}

func (p *QueryJobListResp) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *QueryJobListResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *QueryJobListResp) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *QueryJobListResp) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.LIST, 1)
	listBeginOffset := offset
	offset += thrift.Binary.ListBeginLength()
	var length int
	for _, v := range p.JobList {
		length++
		offset += v.FastWriteNocopy(buf[offset:], w)
	}
	thrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	return offset
}

func (p *QueryJobListResp) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 2)
	offset += thrift.Binary.WriteI64(buf[offset:], p.Total)
	return offset
}

func (p *QueryJobListResp) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.ListBeginLength()
	for _, v := range p.JobList {
		_ = v
		l += v.BLength()
	}
	return l
}

func (p *QueryJobListResp) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *MinerCoreQueryJobListArgs) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MinerCoreQueryJobListArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *MinerCoreQueryJobListArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0
	_field := NewQueryJobListReq()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = _field
	return offset, nil
}

func (p *MinerCoreQueryJobListArgs) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *MinerCoreQueryJobListArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *MinerCoreQueryJobListArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *MinerCoreQueryJobListArgs) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *MinerCoreQueryJobListArgs) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Req.BLength()
	return l
}

func (p *MinerCoreQueryJobListResult) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MinerCoreQueryJobListResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *MinerCoreQueryJobListResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewQueryJobListResp()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

func (p *MinerCoreQueryJobListResult) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *MinerCoreQueryJobListResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *MinerCoreQueryJobListResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *MinerCoreQueryJobListResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *MinerCoreQueryJobListResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *MinerCoreQueryJobListArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *MinerCoreQueryJobListResult) GetResult() interface{} {
	return p.Success
}
