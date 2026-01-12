package revenuecat

import (
	"testing"
)

func TestOverrideOffering(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.OverrideOffering("123", "testUUID")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/offerings/testUUID/override")
}

func TestDeleteOfferingOverride(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.DeleteOfferingOverride("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "DELETE")
	cl.expectPath(t, "/v1/subscribers/123/offerings/override")
}

func TestGetOfferings(t *testing.T) {
	respBody := struct {
		CurrentOfferingID string     `json:"current_offering_id"`
		Offerings         []Offering `json:"offerings"`
	}{
		CurrentOfferingID: "offering_1",
		Offerings: []Offering{
			{
				Description: "offering 1",
				Identifier:  "offering_1",
				Metadata: map[string]interface{}{
					"foo": "bar",
				},
				Packages: []Package{
					{
						Identifier:                "package_1",
						PlatformProductIdentifier: "prod_1",
						PackageType:               MonthlyPackageType,
						Metadata: map[string]interface{}{
							"baz": "qux",
						},
					},
				},
			},
			{
				Description: "offering 2",
				Identifier:  "offering_2",
				Packages: []Package{
					{
						Identifier:                "package_2",
						PlatformProductIdentifier: "prod_2",
						PackageType:               AnnualPackageType,
					},
				},
			},
		},
	}

	cl := newMockClient(t, 200, respBody, nil)
	rc := New("apikey")
	rc.http = cl

	offerings, err := rc.GetOfferings("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if len(offerings.All) != 2 {
		t.Errorf("expected 2 offerings, got %d", len(offerings.All))
	}

	if offerings.Current == nil {
		t.Fatal("expected current offering to be set")
	}

	if offerings.Current.Identifier != "offering_1" {
		t.Errorf("expected current offering to 'offering_1', got %s", offerings.Current.Identifier)
	}

	if offerings.Current.Metadata["foo"] != "bar" {
		t.Errorf("expected metadata 'foo' to be 'bar', got %v", offerings.Current.Metadata["foo"])
	}

	if offerings.Current.Packages[0].Metadata["baz"] != "qux" {
		t.Errorf("expected package metadata 'baz' to be 'qux', got %v", offerings.Current.Packages[0].Metadata["baz"])
	}

	if _, ok := offerings.All["offering_2"]; !ok {
		t.Errorf("expected 'offering_2' to be in All map")
	}

	cl.expectMethod(t, "GET")
	cl.expectPath(t, "/v1/subscribers/123/offerings")
}
