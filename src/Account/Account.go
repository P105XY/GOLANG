package Account import ("fmt")

//DummyType is original struct
type DummyType struct {
    name string value int
}

//InitDummy X aaaaa
func InitDummy(name string, value int) *DummyType {
    dumm := DummyType {
        name:name, value: value;
    }
     return &dumm;
}

//CallingValue X blahblah
func (d DummyType) CallingValue() (string, int) {
    return d.name, d.value
}

//String X blaaahhh
func (d DummyType) String() string {
    return fmt.Sprint("This struct value : ")
}

//method에서, String()으로 함수를 만들고, 반환형으로 string으로 설정.
//fmt.sprint를 반환 시, main에서 fmt.print로 객체 출력을 하면,
//이 method에서 설정한 문자열이 출력된다.
