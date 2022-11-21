package AbstractClothFactory

type SummerClothes struct {
}

// the concrete creator method
func (a SummerClothes) MakeShoe(name string) IShoe {
	return SummerShoe{
		Shoe: Shoe{
			name:  name,
			stype: "summer",
		},
	}
}

// the concrete creator method
func (a SummerClothes) MakeShirt(name string) IShirt {
	return SummerShirt{
		Shirt: Shirt{
			name:  name,
			stype: "summer",
		},
	}
}

func (a SummerClothes) MakePants(name string) IPants {
	return SummerPants{
		Pants: Pants{
			name:  name,
			stype: "summer",
		},
	}
}
