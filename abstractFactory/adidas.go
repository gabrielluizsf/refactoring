package abstractFactory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
    return &AdidasShoe{
        Shoe: Shoe{
            logo: "adidas",
            size: 32,
        },
    }
}

func (a *Adidas) makeShirt() IShirt {
    return &AdidasShirt{
        Shirt: Shirt{
            logo: "adidas",
            size: 39,
        },
    }
}
