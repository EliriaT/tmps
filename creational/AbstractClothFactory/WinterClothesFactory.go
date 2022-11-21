package AbstractClothFactory

type WinterClothes struct {
}

// the concrete creator method
func (a WinterClothes) MakeShoe(name string) IShoe {
	return WinterShoe{
		Shoe: Shoe{
			name:  name,
			stype: "winter",
		},
	}
}

// factory method
func (a WinterClothes) MakeShirt(name string) IShirt {
	return WinterShirt{
		Shirt: Shirt{
			name:  name,
			stype: "winter",
		},
	}
}

// factory method
func (a WinterClothes) MakePants(name string) IPants {
	return WinterPants{
		Pants: Pants{
			name:  name,
			stype: "winter",
		},
	}
}
