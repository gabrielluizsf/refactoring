package abstractFactory

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
    return &NikeShoe{
        Shoe: Shoe{
            logo: "nike",
            size: 39,
        },
    }
}

func (n *Nike) makeShirt() IShirt {
    return &NikeShirt{
        Shirt: Shirt{
            logo: "nike",
            size: 40,
        },
    }
}
