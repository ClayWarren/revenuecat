package revenuecat

// GetOfferings gets the offerings for a specific user.
// https://docs.revenuecat.com/reference#get-offerings
func (c *Client) GetOfferings(userID string) (*Offerings, error) {
	var resp struct {
		CurrentOfferingID string     `json:"current_offering_id"`
		Offerings         []Offering `json:"offerings"`
	}

	err := c.call("GET", "subscribers/"+userID+"/offerings", nil, "", &resp)
	if err != nil {
		return nil, err
	}

	offerings := &Offerings{
		All: make(map[string]Offering),
	}

	for _, o := range resp.Offerings {
		offerings.All[o.Identifier] = o
		if o.Identifier == resp.CurrentOfferingID {
			current := o
			offerings.Current = &current
		}
	}

	return offerings, nil
}

// OverrideOffering overrides the current Offering for a specific user.
// https://docs.revenuecat.com/reference#override-offering
func (c *Client) OverrideOffering(userID string, offeringUUID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("POST", "subscribers/"+userID+"/offerings/"+offeringUUID+"/override", nil, "", &resp)
	return resp.Subscriber, err
}

// DeleteOfferingOverride reset the offering overrides back to the current offering for a specific user.
// https://docs.revenuecat.com/reference#delete-offering-override
func (c *Client) DeleteOfferingOverride(userID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("DELETE", "subscribers/"+userID+"/offerings/override", nil, "", &resp)
	return resp.Subscriber, err
}

// Offerings holds the offerings for a user.
type Offerings struct {
	Current *Offering           `json:"current"`
	All     map[string]Offering `json:"all"`
}

// Offering holds an offering.
type Offering struct {
	Description string                 `json:"description"`
	Identifier  string                 `json:"identifier"`
	Packages    []Package              `json:"packages"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// Package holds a package.
type Package struct {
	Identifier                string                 `json:"identifier"`
	PlatformProductIdentifier string                 `json:"platform_product_identifier"`
	PackageType               PackageType            `json:"package_type"`
	Metadata                  map[string]interface{} `json:"metadata,omitempty"`
}

// PackageType holds the predefined values for a package type.
type PackageType string

// https://docs.revenuecat.com/docs/displaying-products#package-types
const (
	UnknownPackageType    PackageType = "UNKNOWN"
	CustomPackageType     PackageType = "CUSTOM"
	LifetimePackageType   PackageType = "LIFETIME"
	AnnualPackageType     PackageType = "ANNUAL"
	SixMonthPackageType   PackageType = "SIX_MONTH"
	ThreeMonthPackageType PackageType = "THREE_MONTH"
	TwoMonthPackageType   PackageType = "TWO_MONTH"
	MonthlyPackageType    PackageType = "MONTHLY"
	WeeklyPackageType     PackageType = "WEEKLY"
)
