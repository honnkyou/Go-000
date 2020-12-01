package main

import (
	"fmt"
	"os"

	xerrors "github.com/pkg/errors"
)

var errNoRows = errors.New("sql.ErrNoRows")
type query func(string)(string,error)
type item stuct{}

func daoQuery(stmt string) (item, error) {
	if res,err := query(stmt); err != nil{
		return nil, xerrors.Wrapf(errNoRows, "No items were found in %q", stmt)
	}
	return res,nil
	
}

func businessQuery(stmt string) (item, error) {
	return res,err := daoQuery(stmt)
}

func serviceQuery(stmt string) (item, error) {
	res, err := businessQuery(stmt)
	if err != nil{
		return nil, xerrors.WithMessage(err,"added message from service")
	}
	return res,nil
}

func main() {
	res, err := serviceQuery("stmt")
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}
}
