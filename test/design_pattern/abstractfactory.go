
// 抽象工厂 抽象工厂模式与工厂方法模式最大的区别在于，工厂方法模式针对的是一个产品等级结构，而抽象工厂模式则需要面对多个产品等级结构。
// 工厂类接口只有一个create()，抽象工厂有多个
package design_pattern
// 工厂接口和产品接口
type FactoryInterface interface {
	CreatePigMeatBuns() ProductInterface // 创建猪肉馅产品
	Create3SBuns() ProductInterface      // 创建三鲜馅产品
}

type ProductInterface interface {
	Intro()
}
// 实现4种产品
type GDPigMeatBuns struct {
}

func (p GDPigMeatBuns) Intro() {
   fmt.Println("广东猪肉馅包子")
}
.....
// 实现工厂
// 齐市包子铺 
type QSFactory struct {
}

func (qs QSFactory) CreatePigMeatBuns() ProductInterface {
   return QSPigMeatBuns{}
}

func (qs QSFactory) Create3SBuns() ProductInterface {
   return QS3SBuns{}
}
// 广东包子铺
type GDFactory struct {
}

func (gd GDFactory) CreatePigMeatBuns() ProductInterface {
   return GDPigMeatBuns{}
}

func (gd GDFactory) Create3SBuns() ProductInterface {
   return GD3SBuns{}
}

