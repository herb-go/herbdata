package kvdb

import "testing"

func TestNop(t *testing.T) {
	var err error
	var nop Driver = Nop{}
	nop.SetErrorHanlder(nil)
	if nop.Features() != 0 || nop.IsolationLevel() != 0 {
		t.Fatal()
	}
	err = nop.Start()
	if err != nil {
		t.Fatal()
	}
	err = nop.Set(nil, nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.SetWithTTL(nil, nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.SetWithExpired(nil, nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.Get(nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.Delete(nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.Insert(nil, nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.InsertWithTTL(nil, nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.Update(nil, nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.UpdateWithTTL(nil, nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}

	err = nop.SetCounter(nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.SetCounterWithTTL(nil, 0, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.GetCounter(nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.IncreaseCounter(nil, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.IncreaseCounterWithTTL(nil, 0, 0)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.DeleteCounter(nil)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, _, err = nop.Next(nil, 10)
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	_, err = nop.Begin()
	if err != ErrFeatureNotSupported {
		t.Fatal(err)
	}
	err = nop.Stop()
	if err != nil {
		t.Fatal()
	}

}
