package classfile
import "fmt"

type ClassFile struct {
  // magic uint32
  minorVersion uint16
  majorVersion uint16
  constantPool ConstantPool
  accessFlags uint16
  thisClass uint16
  superClass uint16
  interfaces []uint16
  fields []*MemberInfo
  methods []*MemberInfo
  attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
  defer func() {
    if r := recover(); r != nil {
      var ok bool
      err, ok = r.(error)
      if !ok {
        err = fmt.Errorf("%v", r)
      }
    }
  }()

  cr := &ClassReader{classData}
  cf = &ClassFile{}
  cf.read(cr)
  return
}

func (self *ClassFile) read(reader *ClassReader) {
  self.readAndCheckMagic(reader)
  self.readAndCheckVersion(reader)
  self.constantPool = readConstantPool(reader)
  self.accessFlags = reader.readUint16()
  self.thisClass = reader.readUint16()
  self.superClass = reader.readUint16()
  this.interfaces = reader.readUint16()
  self.fields = readMembers(reader, self.constantPool)
  self.methods = readMembers(reader, self.constantPool)
  self.attributes = readAttributes(reader, self.constanPool)
}

func (self *ClassFile) MajorVersion() uint16 {
  return self.majorVersion
}

func (self *ClassFile) MinorVersion() uint16 {
  return self.minorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
  return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
  return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
  return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
  return self.methods
}

func (self *ClassFile) ClassName() string {
  return self.constantPool.getClassName(self.thisClass)
}

// 从常量池查找超类名称, 只有java.lang.Object没有超类
func (self *ClassFile) SuperClassName() string {
  if self.superClass > 0 {
    return self.constantPool.getClassName(self.superClass)
  }
  return ""
}
