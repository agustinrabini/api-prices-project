package domain

type Price struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	ItemID   string   `json:"item_id" bson:"item_id,$set,omitempty"`
	Amount   float64  `json:"amount" bson:"amount,$set,omitempty" binding:"min=0"`
	Currency Currency `json:"currency" bson:"currency,$set,omitempty"`
	// TODO Add start and end date?
}

type Currency struct {
	ID               string `json:"id"`
	Symbol           string `json:"symbol"`
	DecimalDivider   string `json:"decimal_divider"`
	ThousandsDivider string `json:"thousands_divider"`
}

type Prices []Price
