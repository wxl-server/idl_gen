// Code generated by Kitex v0.12.1. DO NOT EDIT.

package fabric_ebl

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *CreateCompanyReq) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetCompanyCode bool = false
	var issetCompanyName bool = false
	var issetCompanyType bool = false
	var issetAdminEmail bool = false
	var issetAdminPassword bool = false
	var issetAdminName bool = false
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
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetCompanyCode = true
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
				issetCompanyName = true
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
				issetCompanyType = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetAdminEmail = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetAdminPassword = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetAdminName = true
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

	if !issetCompanyCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetCompanyName {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetCompanyType {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetAdminEmail {
		fieldId = 4
		goto RequiredFieldNotSetError
	}

	if !issetAdminPassword {
		fieldId = 5
		goto RequiredFieldNotSetError
	}

	if !issetAdminName {
		fieldId = 6
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateCompanyReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_CreateCompanyReq[fieldId]))
}

func (p *CreateCompanyReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.CompanyCode = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.CompanyName = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field CompanyType
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = CompanyType(v)
	}
	p.CompanyType = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.AdminEmail = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastReadField5(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.AdminPassword = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastReadField6(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.AdminName = _field
	return offset, nil
}

func (p *CreateCompanyReq) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *CreateCompanyReq) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *CreateCompanyReq) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *CreateCompanyReq) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.CompanyCode)
	return offset
}

func (p *CreateCompanyReq) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.CompanyName)
	return offset
}

func (p *CreateCompanyReq) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 3)
	offset += thrift.Binary.WriteI32(buf[offset:], int32(p.CompanyType))
	return offset
}

func (p *CreateCompanyReq) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 4)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.AdminEmail)
	return offset
}

func (p *CreateCompanyReq) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 5)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.AdminPassword)
	return offset
}

func (p *CreateCompanyReq) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 6)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.AdminName)
	return offset
}

func (p *CreateCompanyReq) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.CompanyCode)
	return l
}

func (p *CreateCompanyReq) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.CompanyName)
	return l
}

func (p *CreateCompanyReq) field3Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I32Length()
	return l
}

func (p *CreateCompanyReq) field4Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.AdminEmail)
	return l
}

func (p *CreateCompanyReq) field5Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.AdminPassword)
	return l
}

func (p *CreateCompanyReq) field6Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.AdminName)
	return l
}

func (p *CreateCompanyResp) FastRead(buf []byte) (int, error) {

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
		l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
		offset += l
		if err != nil {
			goto SkipFieldError
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *CreateCompanyResp) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *CreateCompanyResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *CreateCompanyResp) BLength() int {
	l := 0
	if p != nil {
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *FabricEblCreateCompanyArgs) FastRead(buf []byte) (int, error) {

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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FabricEblCreateCompanyArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *FabricEblCreateCompanyArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0
	_field := NewCreateCompanyReq()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = _field
	return offset, nil
}

func (p *FabricEblCreateCompanyArgs) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *FabricEblCreateCompanyArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *FabricEblCreateCompanyArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *FabricEblCreateCompanyArgs) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *FabricEblCreateCompanyArgs) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Req.BLength()
	return l
}

func (p *FabricEblCreateCompanyResult) FastRead(buf []byte) (int, error) {

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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FabricEblCreateCompanyResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *FabricEblCreateCompanyResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewCreateCompanyResp()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

func (p *FabricEblCreateCompanyResult) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *FabricEblCreateCompanyResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *FabricEblCreateCompanyResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *FabricEblCreateCompanyResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *FabricEblCreateCompanyResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *FabricEblCreateCompanyArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *FabricEblCreateCompanyResult) GetResult() interface{} {
	return p.Success
}
