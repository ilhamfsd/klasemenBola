package main

const maxTim int = 10
const maxPemain int = 11

type tim struct {
	nama                                                                string
	jumKemenangan, jumGol, jumKebobolan, jumKekalahan, jumSeri, jumskor int
	pemain                                                              tabPemain
}

type pemain struct {
	nama, posisi, np string
}

type tabPemain [maxPemain]pemain

type tabTim [maxTim]tim
