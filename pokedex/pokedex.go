package pokedex

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Pokemon represents a single Pokemon
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"identifier"`
	SpeciesID      int    `json:"species_id"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Order          int    `json:"order"`
	IsDefault      bool   `json:"is_default"`
}

// Pokedex represents a list of Pokemon
type Pokedex struct {
	Pokemon []*Pokemon
}

// FilterBy filters the list of Pokemon in the Pokedex by a given field and match-value
func (p *Pokedex) FilterBy(field string, match interface{}) (*Pokedex, error) {
	fmt.Printf("Finding pokemon with %v = %v (out of %v)\n", field, match, len(p.Pokemon))
	var newList []*Pokemon
	for _, m := range p.Pokemon {
		s := reflect.ValueOf(m).Elem()
		v := s.FieldByName(field)
		if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", match) {
			// fmt.Printf("%#v\n", v)
			newList = append(newList, m)
		}
	}
	p.Pokemon = newList
	return p, nil
}

func (p *Pokedex) String() string {
	indent, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "Error marshaling pokedex"
	}
	return string(indent)
}

// PokedexJSON contains JSON data for all known Pokemon
// last updated 8/17/18
const PokedexJSON = `[
	{
	  "id": 1,
	  "identifier": "bulbasaur",
	  "species_id": 1,
	  "height": 7,
	  "weight": 69,
	  "base_experience": 64,
	  "order": 1,
	  "is_default": true
	},
	{
	  "id": 2,
	  "identifier": "ivysaur",
	  "species_id": 2,
	  "height": 10,
	  "weight": 130,
	  "base_experience": 142,
	  "order": 2,
	  "is_default": true
	},
	{
	  "id": 3,
	  "identifier": "venusaur",
	  "species_id": 3,
	  "height": 20,
	  "weight": 1000,
	  "base_experience": 236,
	  "order": 3,
	  "is_default": true
	},
	{
	  "id": 4,
	  "identifier": "charmander",
	  "species_id": 4,
	  "height": 6,
	  "weight": 85,
	  "base_experience": 62,
	  "order": 5,
	  "is_default": true
	},
	{
	  "id": 5,
	  "identifier": "charmeleon",
	  "species_id": 5,
	  "height": 11,
	  "weight": 190,
	  "base_experience": 142,
	  "order": 6,
	  "is_default": true
	},
	{
	  "id": 6,
	  "identifier": "charizard",
	  "species_id": 6,
	  "height": 17,
	  "weight": 905,
	  "base_experience": 240,
	  "order": 7,
	  "is_default": true
	},
	{
	  "id": 7,
	  "identifier": "squirtle",
	  "species_id": 7,
	  "height": 5,
	  "weight": 90,
	  "base_experience": 63,
	  "order": 10,
	  "is_default": true
	},
	{
	  "id": 8,
	  "identifier": "wartortle",
	  "species_id": 8,
	  "height": 10,
	  "weight": 225,
	  "base_experience": 142,
	  "order": 11,
	  "is_default": true
	},
	{
	  "id": 9,
	  "identifier": "blastoise",
	  "species_id": 9,
	  "height": 16,
	  "weight": 855,
	  "base_experience": 239,
	  "order": 12,
	  "is_default": true
	},
	{
	  "id": 10,
	  "identifier": "caterpie",
	  "species_id": 10,
	  "height": 3,
	  "weight": 29,
	  "base_experience": 39,
	  "order": 14,
	  "is_default": true
	},
	{
	  "id": 11,
	  "identifier": "metapod",
	  "species_id": 11,
	  "height": 7,
	  "weight": 99,
	  "base_experience": 72,
	  "order": 15,
	  "is_default": true
	},
	{
	  "id": 12,
	  "identifier": "butterfree",
	  "species_id": 12,
	  "height": 11,
	  "weight": 320,
	  "base_experience": 178,
	  "order": 16,
	  "is_default": true
	},
	{
	  "id": 13,
	  "identifier": "weedle",
	  "species_id": 13,
	  "height": 3,
	  "weight": 32,
	  "base_experience": 39,
	  "order": 17,
	  "is_default": true
	},
	{
	  "id": 14,
	  "identifier": "kakuna",
	  "species_id": 14,
	  "height": 6,
	  "weight": 100,
	  "base_experience": 72,
	  "order": 18,
	  "is_default": true
	},
	{
	  "id": 15,
	  "identifier": "beedrill",
	  "species_id": 15,
	  "height": 10,
	  "weight": 295,
	  "base_experience": 178,
	  "order": 19,
	  "is_default": true
	},
	{
	  "id": 16,
	  "identifier": "pidgey",
	  "species_id": 16,
	  "height": 3,
	  "weight": 18,
	  "base_experience": 50,
	  "order": 21,
	  "is_default": true
	},
	{
	  "id": 17,
	  "identifier": "pidgeotto",
	  "species_id": 17,
	  "height": 11,
	  "weight": 300,
	  "base_experience": 122,
	  "order": 22,
	  "is_default": true
	},
	{
	  "id": 18,
	  "identifier": "pidgeot",
	  "species_id": 18,
	  "height": 15,
	  "weight": 395,
	  "base_experience": 216,
	  "order": 23,
	  "is_default": true
	},
	{
	  "id": 19,
	  "identifier": "rattata",
	  "species_id": 19,
	  "height": 3,
	  "weight": 35,
	  "base_experience": 51,
	  "order": 25,
	  "is_default": true
	},
	{
	  "id": 20,
	  "identifier": "raticate",
	  "species_id": 20,
	  "height": 7,
	  "weight": 185,
	  "base_experience": 145,
	  "order": 27,
	  "is_default": true
	},
	{
	  "id": 21,
	  "identifier": "spearow",
	  "species_id": 21,
	  "height": 3,
	  "weight": 20,
	  "base_experience": 52,
	  "order": 30,
	  "is_default": true
	},
	{
	  "id": 22,
	  "identifier": "fearow",
	  "species_id": 22,
	  "height": 12,
	  "weight": 380,
	  "base_experience": 155,
	  "order": 31,
	  "is_default": true
	},
	{
	  "id": 23,
	  "identifier": "ekans",
	  "species_id": 23,
	  "height": 20,
	  "weight": 69,
	  "base_experience": 58,
	  "order": 32,
	  "is_default": true
	},
	{
	  "id": 24,
	  "identifier": "arbok",
	  "species_id": 24,
	  "height": 35,
	  "weight": 650,
	  "base_experience": 157,
	  "order": 33,
	  "is_default": true
	},
	{
	  "id": 25,
	  "identifier": "pikachu",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 35,
	  "is_default": true
	},
	{
	  "id": 26,
	  "identifier": "raichu",
	  "species_id": 26,
	  "height": 8,
	  "weight": 300,
	  "base_experience": 218,
	  "order": 43,
	  "is_default": true
	},
	{
	  "id": 27,
	  "identifier": "sandshrew",
	  "species_id": 27,
	  "height": 6,
	  "weight": 120,
	  "base_experience": 60,
	  "order": 45,
	  "is_default": true
	},
	{
	  "id": 28,
	  "identifier": "sandslash",
	  "species_id": 28,
	  "height": 10,
	  "weight": 295,
	  "base_experience": 158,
	  "order": 47,
	  "is_default": true
	},
	{
	  "id": 29,
	  "identifier": "nidoran-f",
	  "species_id": 29,
	  "height": 4,
	  "weight": 70,
	  "base_experience": 55,
	  "order": 49,
	  "is_default": true
	},
	{
	  "id": 30,
	  "identifier": "nidorina",
	  "species_id": 30,
	  "height": 8,
	  "weight": 200,
	  "base_experience": 128,
	  "order": 50,
	  "is_default": true
	},
	{
	  "id": 31,
	  "identifier": "nidoqueen",
	  "species_id": 31,
	  "height": 13,
	  "weight": 600,
	  "base_experience": 227,
	  "order": 51,
	  "is_default": true
	},
	{
	  "id": 32,
	  "identifier": "nidoran-m",
	  "species_id": 32,
	  "height": 5,
	  "weight": 90,
	  "base_experience": 55,
	  "order": 52,
	  "is_default": true
	},
	{
	  "id": 33,
	  "identifier": "nidorino",
	  "species_id": 33,
	  "height": 9,
	  "weight": 195,
	  "base_experience": 128,
	  "order": 53,
	  "is_default": true
	},
	{
	  "id": 34,
	  "identifier": "nidoking",
	  "species_id": 34,
	  "height": 14,
	  "weight": 620,
	  "base_experience": 227,
	  "order": 54,
	  "is_default": true
	},
	{
	  "id": 35,
	  "identifier": "clefairy",
	  "species_id": 35,
	  "height": 6,
	  "weight": 75,
	  "base_experience": 113,
	  "order": 56,
	  "is_default": true
	},
	{
	  "id": 36,
	  "identifier": "clefable",
	  "species_id": 36,
	  "height": 13,
	  "weight": 400,
	  "base_experience": 217,
	  "order": 57,
	  "is_default": true
	},
	{
	  "id": 37,
	  "identifier": "vulpix",
	  "species_id": 37,
	  "height": 6,
	  "weight": 99,
	  "base_experience": 60,
	  "order": 58,
	  "is_default": true
	},
	{
	  "id": 38,
	  "identifier": "ninetales",
	  "species_id": 38,
	  "height": 11,
	  "weight": 199,
	  "base_experience": 177,
	  "order": 60,
	  "is_default": true
	},
	{
	  "id": 39,
	  "identifier": "jigglypuff",
	  "species_id": 39,
	  "height": 5,
	  "weight": 55,
	  "base_experience": 95,
	  "order": 63,
	  "is_default": true
	},
	{
	  "id": 40,
	  "identifier": "wigglytuff",
	  "species_id": 40,
	  "height": 10,
	  "weight": 120,
	  "base_experience": 196,
	  "order": 64,
	  "is_default": true
	},
	{
	  "id": 41,
	  "identifier": "zubat",
	  "species_id": 41,
	  "height": 8,
	  "weight": 75,
	  "base_experience": 49,
	  "order": 65,
	  "is_default": true
	},
	{
	  "id": 42,
	  "identifier": "golbat",
	  "species_id": 42,
	  "height": 16,
	  "weight": 550,
	  "base_experience": 159,
	  "order": 66,
	  "is_default": true
	},
	{
	  "id": 43,
	  "identifier": "oddish",
	  "species_id": 43,
	  "height": 5,
	  "weight": 54,
	  "base_experience": 64,
	  "order": 68,
	  "is_default": true
	},
	{
	  "id": 44,
	  "identifier": "gloom",
	  "species_id": 44,
	  "height": 8,
	  "weight": 86,
	  "base_experience": 138,
	  "order": 69,
	  "is_default": true
	},
	{
	  "id": 45,
	  "identifier": "vileplume",
	  "species_id": 45,
	  "height": 12,
	  "weight": 186,
	  "base_experience": 221,
	  "order": 70,
	  "is_default": true
	},
	{
	  "id": 46,
	  "identifier": "paras",
	  "species_id": 46,
	  "height": 3,
	  "weight": 54,
	  "base_experience": 57,
	  "order": 72,
	  "is_default": true
	},
	{
	  "id": 47,
	  "identifier": "parasect",
	  "species_id": 47,
	  "height": 10,
	  "weight": 295,
	  "base_experience": 142,
	  "order": 73,
	  "is_default": true
	},
	{
	  "id": 48,
	  "identifier": "venonat",
	  "species_id": 48,
	  "height": 10,
	  "weight": 300,
	  "base_experience": 61,
	  "order": 74,
	  "is_default": true
	},
	{
	  "id": 49,
	  "identifier": "venomoth",
	  "species_id": 49,
	  "height": 15,
	  "weight": 125,
	  "base_experience": 158,
	  "order": 75,
	  "is_default": true
	},
	{
	  "id": 50,
	  "identifier": "diglett",
	  "species_id": 50,
	  "height": 2,
	  "weight": 8,
	  "base_experience": 53,
	  "order": 76,
	  "is_default": true
	},
	{
	  "id": 51,
	  "identifier": "dugtrio",
	  "species_id": 51,
	  "height": 7,
	  "weight": 333,
	  "base_experience": 149,
	  "order": 78,
	  "is_default": true
	},
	{
	  "id": 52,
	  "identifier": "meowth",
	  "species_id": 52,
	  "height": 4,
	  "weight": 42,
	  "base_experience": 58,
	  "order": 80,
	  "is_default": true
	},
	{
	  "id": 53,
	  "identifier": "persian",
	  "species_id": 53,
	  "height": 10,
	  "weight": 320,
	  "base_experience": 154,
	  "order": 82,
	  "is_default": true
	},
	{
	  "id": 54,
	  "identifier": "psyduck",
	  "species_id": 54,
	  "height": 8,
	  "weight": 196,
	  "base_experience": 64,
	  "order": 84,
	  "is_default": true
	},
	{
	  "id": 55,
	  "identifier": "golduck",
	  "species_id": 55,
	  "height": 17,
	  "weight": 766,
	  "base_experience": 175,
	  "order": 85,
	  "is_default": true
	},
	{
	  "id": 56,
	  "identifier": "mankey",
	  "species_id": 56,
	  "height": 5,
	  "weight": 280,
	  "base_experience": 61,
	  "order": 86,
	  "is_default": true
	},
	{
	  "id": 57,
	  "identifier": "primeape",
	  "species_id": 57,
	  "height": 10,
	  "weight": 320,
	  "base_experience": 159,
	  "order": 87,
	  "is_default": true
	},
	{
	  "id": 58,
	  "identifier": "growlithe",
	  "species_id": 58,
	  "height": 7,
	  "weight": 190,
	  "base_experience": 70,
	  "order": 88,
	  "is_default": true
	},
	{
	  "id": 59,
	  "identifier": "arcanine",
	  "species_id": 59,
	  "height": 19,
	  "weight": 1550,
	  "base_experience": 194,
	  "order": 89,
	  "is_default": true
	},
	{
	  "id": 60,
	  "identifier": "poliwag",
	  "species_id": 60,
	  "height": 6,
	  "weight": 124,
	  "base_experience": 60,
	  "order": 90,
	  "is_default": true
	},
	{
	  "id": 61,
	  "identifier": "poliwhirl",
	  "species_id": 61,
	  "height": 10,
	  "weight": 200,
	  "base_experience": 135,
	  "order": 91,
	  "is_default": true
	},
	{
	  "id": 62,
	  "identifier": "poliwrath",
	  "species_id": 62,
	  "height": 13,
	  "weight": 540,
	  "base_experience": 230,
	  "order": 92,
	  "is_default": true
	},
	{
	  "id": 63,
	  "identifier": "abra",
	  "species_id": 63,
	  "height": 9,
	  "weight": 195,
	  "base_experience": 62,
	  "order": 94,
	  "is_default": true
	},
	{
	  "id": 64,
	  "identifier": "kadabra",
	  "species_id": 64,
	  "height": 13,
	  "weight": 565,
	  "base_experience": 140,
	  "order": 95,
	  "is_default": true
	},
	{
	  "id": 65,
	  "identifier": "alakazam",
	  "species_id": 65,
	  "height": 15,
	  "weight": 480,
	  "base_experience": 225,
	  "order": 96,
	  "is_default": true
	},
	{
	  "id": 66,
	  "identifier": "machop",
	  "species_id": 66,
	  "height": 8,
	  "weight": 195,
	  "base_experience": 61,
	  "order": 98,
	  "is_default": true
	},
	{
	  "id": 67,
	  "identifier": "machoke",
	  "species_id": 67,
	  "height": 15,
	  "weight": 705,
	  "base_experience": 142,
	  "order": 99,
	  "is_default": true
	},
	{
	  "id": 68,
	  "identifier": "machamp",
	  "species_id": 68,
	  "height": 16,
	  "weight": 1300,
	  "base_experience": 227,
	  "order": 100,
	  "is_default": true
	},
	{
	  "id": 69,
	  "identifier": "bellsprout",
	  "species_id": 69,
	  "height": 7,
	  "weight": 40,
	  "base_experience": 60,
	  "order": 101,
	  "is_default": true
	},
	{
	  "id": 70,
	  "identifier": "weepinbell",
	  "species_id": 70,
	  "height": 10,
	  "weight": 64,
	  "base_experience": 137,
	  "order": 102,
	  "is_default": true
	},
	{
	  "id": 71,
	  "identifier": "victreebel",
	  "species_id": 71,
	  "height": 17,
	  "weight": 155,
	  "base_experience": 221,
	  "order": 103,
	  "is_default": true
	},
	{
	  "id": 72,
	  "identifier": "tentacool",
	  "species_id": 72,
	  "height": 9,
	  "weight": 455,
	  "base_experience": 67,
	  "order": 104,
	  "is_default": true
	},
	{
	  "id": 73,
	  "identifier": "tentacruel",
	  "species_id": 73,
	  "height": 16,
	  "weight": 550,
	  "base_experience": 180,
	  "order": 105,
	  "is_default": true
	},
	{
	  "id": 74,
	  "identifier": "geodude",
	  "species_id": 74,
	  "height": 4,
	  "weight": 200,
	  "base_experience": 60,
	  "order": 106,
	  "is_default": true
	},
	{
	  "id": 75,
	  "identifier": "graveler",
	  "species_id": 75,
	  "height": 10,
	  "weight": 1050,
	  "base_experience": 137,
	  "order": 108,
	  "is_default": true
	},
	{
	  "id": 76,
	  "identifier": "golem",
	  "species_id": 76,
	  "height": 14,
	  "weight": 3000,
	  "base_experience": 223,
	  "order": 110,
	  "is_default": true
	},
	{
	  "id": 77,
	  "identifier": "ponyta",
	  "species_id": 77,
	  "height": 10,
	  "weight": 300,
	  "base_experience": 82,
	  "order": 112,
	  "is_default": true
	},
	{
	  "id": 78,
	  "identifier": "rapidash",
	  "species_id": 78,
	  "height": 17,
	  "weight": 950,
	  "base_experience": 175,
	  "order": 113,
	  "is_default": true
	},
	{
	  "id": 79,
	  "identifier": "slowpoke",
	  "species_id": 79,
	  "height": 12,
	  "weight": 360,
	  "base_experience": 63,
	  "order": 114,
	  "is_default": true
	},
	{
	  "id": 80,
	  "identifier": "slowbro",
	  "species_id": 80,
	  "height": 16,
	  "weight": 785,
	  "base_experience": 172,
	  "order": 115,
	  "is_default": true
	},
	{
	  "id": 81,
	  "identifier": "magnemite",
	  "species_id": 81,
	  "height": 3,
	  "weight": 60,
	  "base_experience": 65,
	  "order": 118,
	  "is_default": true
	},
	{
	  "id": 82,
	  "identifier": "magneton",
	  "species_id": 82,
	  "height": 10,
	  "weight": 600,
	  "base_experience": 163,
	  "order": 119,
	  "is_default": true
	},
	{
	  "id": 83,
	  "identifier": "farfetchd",
	  "species_id": 83,
	  "height": 8,
	  "weight": 150,
	  "base_experience": 132,
	  "order": 121,
	  "is_default": true
	},
	{
	  "id": 84,
	  "identifier": "doduo",
	  "species_id": 84,
	  "height": 14,
	  "weight": 392,
	  "base_experience": 62,
	  "order": 122,
	  "is_default": true
	},
	{
	  "id": 85,
	  "identifier": "dodrio",
	  "species_id": 85,
	  "height": 18,
	  "weight": 852,
	  "base_experience": 165,
	  "order": 123,
	  "is_default": true
	},
	{
	  "id": 86,
	  "identifier": "seel",
	  "species_id": 86,
	  "height": 11,
	  "weight": 900,
	  "base_experience": 65,
	  "order": 124,
	  "is_default": true
	},
	{
	  "id": 87,
	  "identifier": "dewgong",
	  "species_id": 87,
	  "height": 17,
	  "weight": 1200,
	  "base_experience": 166,
	  "order": 125,
	  "is_default": true
	},
	{
	  "id": 88,
	  "identifier": "grimer",
	  "species_id": 88,
	  "height": 9,
	  "weight": 300,
	  "base_experience": 65,
	  "order": 126,
	  "is_default": true
	},
	{
	  "id": 89,
	  "identifier": "muk",
	  "species_id": 89,
	  "height": 12,
	  "weight": 300,
	  "base_experience": 175,
	  "order": 128,
	  "is_default": true
	},
	{
	  "id": 90,
	  "identifier": "shellder",
	  "species_id": 90,
	  "height": 3,
	  "weight": 40,
	  "base_experience": 61,
	  "order": 130,
	  "is_default": true
	},
	{
	  "id": 91,
	  "identifier": "cloyster",
	  "species_id": 91,
	  "height": 15,
	  "weight": 1325,
	  "base_experience": 184,
	  "order": 131,
	  "is_default": true
	},
	{
	  "id": 92,
	  "identifier": "gastly",
	  "species_id": 92,
	  "height": 13,
	  "weight": 1,
	  "base_experience": 62,
	  "order": 132,
	  "is_default": true
	},
	{
	  "id": 93,
	  "identifier": "haunter",
	  "species_id": 93,
	  "height": 16,
	  "weight": 1,
	  "base_experience": 142,
	  "order": 133,
	  "is_default": true
	},
	{
	  "id": 94,
	  "identifier": "gengar",
	  "species_id": 94,
	  "height": 15,
	  "weight": 405,
	  "base_experience": 225,
	  "order": 134,
	  "is_default": true
	},
	{
	  "id": 95,
	  "identifier": "onix",
	  "species_id": 95,
	  "height": 88,
	  "weight": 2100,
	  "base_experience": 77,
	  "order": 136,
	  "is_default": true
	},
	{
	  "id": 96,
	  "identifier": "drowzee",
	  "species_id": 96,
	  "height": 10,
	  "weight": 324,
	  "base_experience": 66,
	  "order": 139,
	  "is_default": true
	},
	{
	  "id": 97,
	  "identifier": "hypno",
	  "species_id": 97,
	  "height": 16,
	  "weight": 756,
	  "base_experience": 169,
	  "order": 140,
	  "is_default": true
	},
	{
	  "id": 98,
	  "identifier": "krabby",
	  "species_id": 98,
	  "height": 4,
	  "weight": 65,
	  "base_experience": 65,
	  "order": 141,
	  "is_default": true
	},
	{
	  "id": 99,
	  "identifier": "kingler",
	  "species_id": 99,
	  "height": 13,
	  "weight": 600,
	  "base_experience": 166,
	  "order": 142,
	  "is_default": true
	},
	{
	  "id": 100,
	  "identifier": "voltorb",
	  "species_id": 100,
	  "height": 5,
	  "weight": 104,
	  "base_experience": 66,
	  "order": 143,
	  "is_default": true
	},
	{
	  "id": 101,
	  "identifier": "electrode",
	  "species_id": 101,
	  "height": 12,
	  "weight": 666,
	  "base_experience": 172,
	  "order": 144,
	  "is_default": true
	},
	{
	  "id": 102,
	  "identifier": "exeggcute",
	  "species_id": 102,
	  "height": 4,
	  "weight": 25,
	  "base_experience": 65,
	  "order": 145,
	  "is_default": true
	},
	{
	  "id": 103,
	  "identifier": "exeggutor",
	  "species_id": 103,
	  "height": 20,
	  "weight": 1200,
	  "base_experience": 186,
	  "order": 146,
	  "is_default": true
	},
	{
	  "id": 104,
	  "identifier": "cubone",
	  "species_id": 104,
	  "height": 4,
	  "weight": 65,
	  "base_experience": 64,
	  "order": 148,
	  "is_default": true
	},
	{
	  "id": 105,
	  "identifier": "marowak",
	  "species_id": 105,
	  "height": 10,
	  "weight": 450,
	  "base_experience": 149,
	  "order": 149,
	  "is_default": true
	},
	{
	  "id": 106,
	  "identifier": "hitmonlee",
	  "species_id": 106,
	  "height": 15,
	  "weight": 498,
	  "base_experience": 159,
	  "order": 153,
	  "is_default": true
	},
	{
	  "id": 107,
	  "identifier": "hitmonchan",
	  "species_id": 107,
	  "height": 14,
	  "weight": 502,
	  "base_experience": 159,
	  "order": 154,
	  "is_default": true
	},
	{
	  "id": 108,
	  "identifier": "lickitung",
	  "species_id": 108,
	  "height": 12,
	  "weight": 655,
	  "base_experience": 77,
	  "order": 156,
	  "is_default": true
	},
	{
	  "id": 109,
	  "identifier": "koffing",
	  "species_id": 109,
	  "height": 6,
	  "weight": 10,
	  "base_experience": 68,
	  "order": 158,
	  "is_default": true
	},
	{
	  "id": 110,
	  "identifier": "weezing",
	  "species_id": 110,
	  "height": 12,
	  "weight": 95,
	  "base_experience": 172,
	  "order": 159,
	  "is_default": true
	},
	{
	  "id": 111,
	  "identifier": "rhyhorn",
	  "species_id": 111,
	  "height": 10,
	  "weight": 1150,
	  "base_experience": 69,
	  "order": 160,
	  "is_default": true
	},
	{
	  "id": 112,
	  "identifier": "rhydon",
	  "species_id": 112,
	  "height": 19,
	  "weight": 1200,
	  "base_experience": 170,
	  "order": 161,
	  "is_default": true
	},
	{
	  "id": 113,
	  "identifier": "chansey",
	  "species_id": 113,
	  "height": 11,
	  "weight": 346,
	  "base_experience": 395,
	  "order": 164,
	  "is_default": true
	},
	{
	  "id": 114,
	  "identifier": "tangela",
	  "species_id": 114,
	  "height": 10,
	  "weight": 350,
	  "base_experience": 87,
	  "order": 166,
	  "is_default": true
	},
	{
	  "id": 115,
	  "identifier": "kangaskhan",
	  "species_id": 115,
	  "height": 22,
	  "weight": 800,
	  "base_experience": 172,
	  "order": 168,
	  "is_default": true
	},
	{
	  "id": 116,
	  "identifier": "horsea",
	  "species_id": 116,
	  "height": 4,
	  "weight": 80,
	  "base_experience": 59,
	  "order": 170,
	  "is_default": true
	},
	{
	  "id": 117,
	  "identifier": "seadra",
	  "species_id": 117,
	  "height": 12,
	  "weight": 250,
	  "base_experience": 154,
	  "order": 171,
	  "is_default": true
	},
	{
	  "id": 118,
	  "identifier": "goldeen",
	  "species_id": 118,
	  "height": 6,
	  "weight": 150,
	  "base_experience": 64,
	  "order": 173,
	  "is_default": true
	},
	{
	  "id": 119,
	  "identifier": "seaking",
	  "species_id": 119,
	  "height": 13,
	  "weight": 390,
	  "base_experience": 158,
	  "order": 174,
	  "is_default": true
	},
	{
	  "id": 120,
	  "identifier": "staryu",
	  "species_id": 120,
	  "height": 8,
	  "weight": 345,
	  "base_experience": 68,
	  "order": 175,
	  "is_default": true
	},
	{
	  "id": 121,
	  "identifier": "starmie",
	  "species_id": 121,
	  "height": 11,
	  "weight": 800,
	  "base_experience": 182,
	  "order": 176,
	  "is_default": true
	},
	{
	  "id": 122,
	  "identifier": "mr-mime",
	  "species_id": 122,
	  "height": 13,
	  "weight": 545,
	  "base_experience": 161,
	  "order": 178,
	  "is_default": true
	},
	{
	  "id": 123,
	  "identifier": "scyther",
	  "species_id": 123,
	  "height": 15,
	  "weight": 560,
	  "base_experience": 100,
	  "order": 179,
	  "is_default": true
	},
	{
	  "id": 124,
	  "identifier": "jynx",
	  "species_id": 124,
	  "height": 14,
	  "weight": 406,
	  "base_experience": 159,
	  "order": 183,
	  "is_default": true
	},
	{
	  "id": 125,
	  "identifier": "electabuzz",
	  "species_id": 125,
	  "height": 11,
	  "weight": 300,
	  "base_experience": 172,
	  "order": 185,
	  "is_default": true
	},
	{
	  "id": 126,
	  "identifier": "magmar",
	  "species_id": 126,
	  "height": 13,
	  "weight": 445,
	  "base_experience": 173,
	  "order": 188,
	  "is_default": true
	},
	{
	  "id": 127,
	  "identifier": "pinsir",
	  "species_id": 127,
	  "height": 15,
	  "weight": 550,
	  "base_experience": 175,
	  "order": 190,
	  "is_default": true
	},
	{
	  "id": 128,
	  "identifier": "tauros",
	  "species_id": 128,
	  "height": 14,
	  "weight": 884,
	  "base_experience": 172,
	  "order": 192,
	  "is_default": true
	},
	{
	  "id": 129,
	  "identifier": "magikarp",
	  "species_id": 129,
	  "height": 9,
	  "weight": 100,
	  "base_experience": 40,
	  "order": 193,
	  "is_default": true
	},
	{
	  "id": 130,
	  "identifier": "gyarados",
	  "species_id": 130,
	  "height": 65,
	  "weight": 2350,
	  "base_experience": 189,
	  "order": 194,
	  "is_default": true
	},
	{
	  "id": 131,
	  "identifier": "lapras",
	  "species_id": 131,
	  "height": 25,
	  "weight": 2200,
	  "base_experience": 187,
	  "order": 196,
	  "is_default": true
	},
	{
	  "id": 132,
	  "identifier": "ditto",
	  "species_id": 132,
	  "height": 3,
	  "weight": 40,
	  "base_experience": 101,
	  "order": 197,
	  "is_default": true
	},
	{
	  "id": 133,
	  "identifier": "eevee",
	  "species_id": 133,
	  "height": 3,
	  "weight": 65,
	  "base_experience": 65,
	  "order": 198,
	  "is_default": true
	},
	{
	  "id": 134,
	  "identifier": "vaporeon",
	  "species_id": 134,
	  "height": 10,
	  "weight": 290,
	  "base_experience": 184,
	  "order": 199,
	  "is_default": true
	},
	{
	  "id": 135,
	  "identifier": "jolteon",
	  "species_id": 135,
	  "height": 8,
	  "weight": 245,
	  "base_experience": 184,
	  "order": 200,
	  "is_default": true
	},
	{
	  "id": 136,
	  "identifier": "flareon",
	  "species_id": 136,
	  "height": 9,
	  "weight": 250,
	  "base_experience": 184,
	  "order": 201,
	  "is_default": true
	},
	{
	  "id": 137,
	  "identifier": "porygon",
	  "species_id": 137,
	  "height": 8,
	  "weight": 365,
	  "base_experience": 79,
	  "order": 207,
	  "is_default": true
	},
	{
	  "id": 138,
	  "identifier": "omanyte",
	  "species_id": 138,
	  "height": 4,
	  "weight": 75,
	  "base_experience": 71,
	  "order": 210,
	  "is_default": true
	},
	{
	  "id": 139,
	  "identifier": "omastar",
	  "species_id": 139,
	  "height": 10,
	  "weight": 350,
	  "base_experience": 173,
	  "order": 211,
	  "is_default": true
	},
	{
	  "id": 140,
	  "identifier": "kabuto",
	  "species_id": 140,
	  "height": 5,
	  "weight": 115,
	  "base_experience": 71,
	  "order": 212,
	  "is_default": true
	},
	{
	  "id": 141,
	  "identifier": "kabutops",
	  "species_id": 141,
	  "height": 13,
	  "weight": 405,
	  "base_experience": 173,
	  "order": 213,
	  "is_default": true
	},
	{
	  "id": 142,
	  "identifier": "aerodactyl",
	  "species_id": 142,
	  "height": 18,
	  "weight": 590,
	  "base_experience": 180,
	  "order": 214,
	  "is_default": true
	},
	{
	  "id": 143,
	  "identifier": "snorlax",
	  "species_id": 143,
	  "height": 21,
	  "weight": 4600,
	  "base_experience": 189,
	  "order": 217,
	  "is_default": true
	},
	{
	  "id": 144,
	  "identifier": "articuno",
	  "species_id": 144,
	  "height": 17,
	  "weight": 554,
	  "base_experience": 261,
	  "order": 218,
	  "is_default": true
	},
	{
	  "id": 145,
	  "identifier": "zapdos",
	  "species_id": 145,
	  "height": 16,
	  "weight": 526,
	  "base_experience": 261,
	  "order": 219,
	  "is_default": true
	},
	{
	  "id": 146,
	  "identifier": "moltres",
	  "species_id": 146,
	  "height": 20,
	  "weight": 600,
	  "base_experience": 261,
	  "order": 220,
	  "is_default": true
	},
	{
	  "id": 147,
	  "identifier": "dratini",
	  "species_id": 147,
	  "height": 18,
	  "weight": 33,
	  "base_experience": 60,
	  "order": 221,
	  "is_default": true
	},
	{
	  "id": 148,
	  "identifier": "dragonair",
	  "species_id": 148,
	  "height": 40,
	  "weight": 165,
	  "base_experience": 147,
	  "order": 222,
	  "is_default": true
	},
	{
	  "id": 149,
	  "identifier": "dragonite",
	  "species_id": 149,
	  "height": 22,
	  "weight": 2100,
	  "base_experience": 270,
	  "order": 223,
	  "is_default": true
	},
	{
	  "id": 150,
	  "identifier": "mewtwo",
	  "species_id": 150,
	  "height": 20,
	  "weight": 1220,
	  "base_experience": 306,
	  "order": 224,
	  "is_default": true
	},
	{
	  "id": 151,
	  "identifier": "mew",
	  "species_id": 151,
	  "height": 4,
	  "weight": 40,
	  "base_experience": 270,
	  "order": 227,
	  "is_default": true
	},
	{
	  "id": 152,
	  "identifier": "chikorita",
	  "species_id": 152,
	  "height": 9,
	  "weight": 64,
	  "base_experience": 64,
	  "order": 228,
	  "is_default": true
	},
	{
	  "id": 153,
	  "identifier": "bayleef",
	  "species_id": 153,
	  "height": 12,
	  "weight": 158,
	  "base_experience": 142,
	  "order": 229,
	  "is_default": true
	},
	{
	  "id": 154,
	  "identifier": "meganium",
	  "species_id": 154,
	  "height": 18,
	  "weight": 1005,
	  "base_experience": 236,
	  "order": 230,
	  "is_default": true
	},
	{
	  "id": 155,
	  "identifier": "cyndaquil",
	  "species_id": 155,
	  "height": 5,
	  "weight": 79,
	  "base_experience": 62,
	  "order": 231,
	  "is_default": true
	},
	{
	  "id": 156,
	  "identifier": "quilava",
	  "species_id": 156,
	  "height": 9,
	  "weight": 190,
	  "base_experience": 142,
	  "order": 232,
	  "is_default": true
	},
	{
	  "id": 157,
	  "identifier": "typhlosion",
	  "species_id": 157,
	  "height": 17,
	  "weight": 795,
	  "base_experience": 240,
	  "order": 233,
	  "is_default": true
	},
	{
	  "id": 158,
	  "identifier": "totodile",
	  "species_id": 158,
	  "height": 6,
	  "weight": 95,
	  "base_experience": 63,
	  "order": 234,
	  "is_default": true
	},
	{
	  "id": 159,
	  "identifier": "croconaw",
	  "species_id": 159,
	  "height": 11,
	  "weight": 250,
	  "base_experience": 142,
	  "order": 235,
	  "is_default": true
	},
	{
	  "id": 160,
	  "identifier": "feraligatr",
	  "species_id": 160,
	  "height": 23,
	  "weight": 888,
	  "base_experience": 239,
	  "order": 236,
	  "is_default": true
	},
	{
	  "id": 161,
	  "identifier": "sentret",
	  "species_id": 161,
	  "height": 8,
	  "weight": 60,
	  "base_experience": 43,
	  "order": 237,
	  "is_default": true
	},
	{
	  "id": 162,
	  "identifier": "furret",
	  "species_id": 162,
	  "height": 18,
	  "weight": 325,
	  "base_experience": 145,
	  "order": 238,
	  "is_default": true
	},
	{
	  "id": 163,
	  "identifier": "hoothoot",
	  "species_id": 163,
	  "height": 7,
	  "weight": 212,
	  "base_experience": 52,
	  "order": 239,
	  "is_default": true
	},
	{
	  "id": 164,
	  "identifier": "noctowl",
	  "species_id": 164,
	  "height": 16,
	  "weight": 408,
	  "base_experience": 158,
	  "order": 240,
	  "is_default": true
	},
	{
	  "id": 165,
	  "identifier": "ledyba",
	  "species_id": 165,
	  "height": 10,
	  "weight": 108,
	  "base_experience": 53,
	  "order": 241,
	  "is_default": true
	},
	{
	  "id": 166,
	  "identifier": "ledian",
	  "species_id": 166,
	  "height": 14,
	  "weight": 356,
	  "base_experience": 137,
	  "order": 242,
	  "is_default": true
	},
	{
	  "id": 167,
	  "identifier": "spinarak",
	  "species_id": 167,
	  "height": 5,
	  "weight": 85,
	  "base_experience": 50,
	  "order": 243,
	  "is_default": true
	},
	{
	  "id": 168,
	  "identifier": "ariados",
	  "species_id": 168,
	  "height": 11,
	  "weight": 335,
	  "base_experience": 140,
	  "order": 244,
	  "is_default": true
	},
	{
	  "id": 169,
	  "identifier": "crobat",
	  "species_id": 169,
	  "height": 18,
	  "weight": 750,
	  "base_experience": 241,
	  "order": 67,
	  "is_default": true
	},
	{
	  "id": 170,
	  "identifier": "chinchou",
	  "species_id": 170,
	  "height": 5,
	  "weight": 120,
	  "base_experience": 66,
	  "order": 245,
	  "is_default": true
	},
	{
	  "id": 171,
	  "identifier": "lanturn",
	  "species_id": 171,
	  "height": 12,
	  "weight": 225,
	  "base_experience": 161,
	  "order": 246,
	  "is_default": true
	},
	{
	  "id": 172,
	  "identifier": "pichu",
	  "species_id": 172,
	  "height": 3,
	  "weight": 20,
	  "base_experience": 41,
	  "order": 34,
	  "is_default": true
	},
	{
	  "id": 173,
	  "identifier": "cleffa",
	  "species_id": 173,
	  "height": 3,
	  "weight": 30,
	  "base_experience": 44,
	  "order": 55,
	  "is_default": true
	},
	{
	  "id": 174,
	  "identifier": "igglybuff",
	  "species_id": 174,
	  "height": 3,
	  "weight": 10,
	  "base_experience": 42,
	  "order": 62,
	  "is_default": true
	},
	{
	  "id": 175,
	  "identifier": "togepi",
	  "species_id": 175,
	  "height": 3,
	  "weight": 15,
	  "base_experience": 49,
	  "order": 247,
	  "is_default": true
	},
	{
	  "id": 176,
	  "identifier": "togetic",
	  "species_id": 176,
	  "height": 6,
	  "weight": 32,
	  "base_experience": 142,
	  "order": 248,
	  "is_default": true
	},
	{
	  "id": 177,
	  "identifier": "natu",
	  "species_id": 177,
	  "height": 2,
	  "weight": 20,
	  "base_experience": 64,
	  "order": 250,
	  "is_default": true
	},
	{
	  "id": 178,
	  "identifier": "xatu",
	  "species_id": 178,
	  "height": 15,
	  "weight": 150,
	  "base_experience": 165,
	  "order": 251,
	  "is_default": true
	},
	{
	  "id": 179,
	  "identifier": "mareep",
	  "species_id": 179,
	  "height": 6,
	  "weight": 78,
	  "base_experience": 56,
	  "order": 252,
	  "is_default": true
	},
	{
	  "id": 180,
	  "identifier": "flaaffy",
	  "species_id": 180,
	  "height": 8,
	  "weight": 133,
	  "base_experience": 128,
	  "order": 253,
	  "is_default": true
	},
	{
	  "id": 181,
	  "identifier": "ampharos",
	  "species_id": 181,
	  "height": 14,
	  "weight": 615,
	  "base_experience": 230,
	  "order": 254,
	  "is_default": true
	},
	{
	  "id": 182,
	  "identifier": "bellossom",
	  "species_id": 182,
	  "height": 4,
	  "weight": 58,
	  "base_experience": 221,
	  "order": 71,
	  "is_default": true
	},
	{
	  "id": 183,
	  "identifier": "marill",
	  "species_id": 183,
	  "height": 4,
	  "weight": 85,
	  "base_experience": 88,
	  "order": 257,
	  "is_default": true
	},
	{
	  "id": 184,
	  "identifier": "azumarill",
	  "species_id": 184,
	  "height": 8,
	  "weight": 285,
	  "base_experience": 189,
	  "order": 258,
	  "is_default": true
	},
	{
	  "id": 185,
	  "identifier": "sudowoodo",
	  "species_id": 185,
	  "height": 12,
	  "weight": 380,
	  "base_experience": 144,
	  "order": 260,
	  "is_default": true
	},
	{
	  "id": 186,
	  "identifier": "politoed",
	  "species_id": 186,
	  "height": 11,
	  "weight": 339,
	  "base_experience": 225,
	  "order": 93,
	  "is_default": true
	},
	{
	  "id": 187,
	  "identifier": "hoppip",
	  "species_id": 187,
	  "height": 4,
	  "weight": 5,
	  "base_experience": 50,
	  "order": 261,
	  "is_default": true
	},
	{
	  "id": 188,
	  "identifier": "skiploom",
	  "species_id": 188,
	  "height": 6,
	  "weight": 10,
	  "base_experience": 119,
	  "order": 262,
	  "is_default": true
	},
	{
	  "id": 189,
	  "identifier": "jumpluff",
	  "species_id": 189,
	  "height": 8,
	  "weight": 30,
	  "base_experience": 207,
	  "order": 263,
	  "is_default": true
	},
	{
	  "id": 190,
	  "identifier": "aipom",
	  "species_id": 190,
	  "height": 8,
	  "weight": 115,
	  "base_experience": 72,
	  "order": 264,
	  "is_default": true
	},
	{
	  "id": 191,
	  "identifier": "sunkern",
	  "species_id": 191,
	  "height": 3,
	  "weight": 18,
	  "base_experience": 36,
	  "order": 266,
	  "is_default": true
	},
	{
	  "id": 192,
	  "identifier": "sunflora",
	  "species_id": 192,
	  "height": 8,
	  "weight": 85,
	  "base_experience": 149,
	  "order": 267,
	  "is_default": true
	},
	{
	  "id": 193,
	  "identifier": "yanma",
	  "species_id": 193,
	  "height": 12,
	  "weight": 380,
	  "base_experience": 78,
	  "order": 268,
	  "is_default": true
	},
	{
	  "id": 194,
	  "identifier": "wooper",
	  "species_id": 194,
	  "height": 4,
	  "weight": 85,
	  "base_experience": 42,
	  "order": 270,
	  "is_default": true
	},
	{
	  "id": 195,
	  "identifier": "quagsire",
	  "species_id": 195,
	  "height": 14,
	  "weight": 750,
	  "base_experience": 151,
	  "order": 271,
	  "is_default": true
	},
	{
	  "id": 196,
	  "identifier": "espeon",
	  "species_id": 196,
	  "height": 9,
	  "weight": 265,
	  "base_experience": 184,
	  "order": 202,
	  "is_default": true
	},
	{
	  "id": 197,
	  "identifier": "umbreon",
	  "species_id": 197,
	  "height": 10,
	  "weight": 270,
	  "base_experience": 184,
	  "order": 203,
	  "is_default": true
	},
	{
	  "id": 198,
	  "identifier": "murkrow",
	  "species_id": 198,
	  "height": 5,
	  "weight": 21,
	  "base_experience": 81,
	  "order": 272,
	  "is_default": true
	},
	{
	  "id": 199,
	  "identifier": "slowking",
	  "species_id": 199,
	  "height": 20,
	  "weight": 795,
	  "base_experience": 172,
	  "order": 117,
	  "is_default": true
	},
	{
	  "id": 200,
	  "identifier": "misdreavus",
	  "species_id": 200,
	  "height": 7,
	  "weight": 10,
	  "base_experience": 87,
	  "order": 274,
	  "is_default": true
	},
	{
	  "id": 201,
	  "identifier": "unown",
	  "species_id": 201,
	  "height": 5,
	  "weight": 50,
	  "base_experience": 118,
	  "order": 276,
	  "is_default": true
	},
	{
	  "id": 202,
	  "identifier": "wobbuffet",
	  "species_id": 202,
	  "height": 13,
	  "weight": 285,
	  "base_experience": 142,
	  "order": 278,
	  "is_default": true
	},
	{
	  "id": 203,
	  "identifier": "girafarig",
	  "species_id": 203,
	  "height": 15,
	  "weight": 415,
	  "base_experience": 159,
	  "order": 279,
	  "is_default": true
	},
	{
	  "id": 204,
	  "identifier": "pineco",
	  "species_id": 204,
	  "height": 6,
	  "weight": 72,
	  "base_experience": 58,
	  "order": 280,
	  "is_default": true
	},
	{
	  "id": 205,
	  "identifier": "forretress",
	  "species_id": 205,
	  "height": 12,
	  "weight": 1258,
	  "base_experience": 163,
	  "order": 281,
	  "is_default": true
	},
	{
	  "id": 206,
	  "identifier": "dunsparce",
	  "species_id": 206,
	  "height": 15,
	  "weight": 140,
	  "base_experience": 145,
	  "order": 282,
	  "is_default": true
	},
	{
	  "id": 207,
	  "identifier": "gligar",
	  "species_id": 207,
	  "height": 11,
	  "weight": 648,
	  "base_experience": 86,
	  "order": 283,
	  "is_default": true
	},
	{
	  "id": 208,
	  "identifier": "steelix",
	  "species_id": 208,
	  "height": 92,
	  "weight": 4000,
	  "base_experience": 179,
	  "order": 137,
	  "is_default": true
	},
	{
	  "id": 209,
	  "identifier": "snubbull",
	  "species_id": 209,
	  "height": 6,
	  "weight": 78,
	  "base_experience": 60,
	  "order": 285,
	  "is_default": true
	},
	{
	  "id": 210,
	  "identifier": "granbull",
	  "species_id": 210,
	  "height": 14,
	  "weight": 487,
	  "base_experience": 158,
	  "order": 286,
	  "is_default": true
	},
	{
	  "id": 211,
	  "identifier": "qwilfish",
	  "species_id": 211,
	  "height": 5,
	  "weight": 39,
	  "base_experience": 88,
	  "order": 287,
	  "is_default": true
	},
	{
	  "id": 212,
	  "identifier": "scizor",
	  "species_id": 212,
	  "height": 18,
	  "weight": 1180,
	  "base_experience": 175,
	  "order": 180,
	  "is_default": true
	},
	{
	  "id": 213,
	  "identifier": "shuckle",
	  "species_id": 213,
	  "height": 6,
	  "weight": 205,
	  "base_experience": 177,
	  "order": 288,
	  "is_default": true
	},
	{
	  "id": 214,
	  "identifier": "heracross",
	  "species_id": 214,
	  "height": 15,
	  "weight": 540,
	  "base_experience": 175,
	  "order": 289,
	  "is_default": true
	},
	{
	  "id": 215,
	  "identifier": "sneasel",
	  "species_id": 215,
	  "height": 9,
	  "weight": 280,
	  "base_experience": 86,
	  "order": 291,
	  "is_default": true
	},
	{
	  "id": 216,
	  "identifier": "teddiursa",
	  "species_id": 216,
	  "height": 6,
	  "weight": 88,
	  "base_experience": 66,
	  "order": 293,
	  "is_default": true
	},
	{
	  "id": 217,
	  "identifier": "ursaring",
	  "species_id": 217,
	  "height": 18,
	  "weight": 1258,
	  "base_experience": 175,
	  "order": 294,
	  "is_default": true
	},
	{
	  "id": 218,
	  "identifier": "slugma",
	  "species_id": 218,
	  "height": 7,
	  "weight": 350,
	  "base_experience": 50,
	  "order": 295,
	  "is_default": true
	},
	{
	  "id": 219,
	  "identifier": "magcargo",
	  "species_id": 219,
	  "height": 8,
	  "weight": 550,
	  "base_experience": 151,
	  "order": 296,
	  "is_default": true
	},
	{
	  "id": 220,
	  "identifier": "swinub",
	  "species_id": 220,
	  "height": 4,
	  "weight": 65,
	  "base_experience": 50,
	  "order": 297,
	  "is_default": true
	},
	{
	  "id": 221,
	  "identifier": "piloswine",
	  "species_id": 221,
	  "height": 11,
	  "weight": 558,
	  "base_experience": 158,
	  "order": 298,
	  "is_default": true
	},
	{
	  "id": 222,
	  "identifier": "corsola",
	  "species_id": 222,
	  "height": 6,
	  "weight": 50,
	  "base_experience": 144,
	  "order": 300,
	  "is_default": true
	},
	{
	  "id": 223,
	  "identifier": "remoraid",
	  "species_id": 223,
	  "height": 6,
	  "weight": 120,
	  "base_experience": 60,
	  "order": 301,
	  "is_default": true
	},
	{
	  "id": 224,
	  "identifier": "octillery",
	  "species_id": 224,
	  "height": 9,
	  "weight": 285,
	  "base_experience": 168,
	  "order": 302,
	  "is_default": true
	},
	{
	  "id": 225,
	  "identifier": "delibird",
	  "species_id": 225,
	  "height": 9,
	  "weight": 160,
	  "base_experience": 116,
	  "order": 303,
	  "is_default": true
	},
	{
	  "id": 226,
	  "identifier": "mantine",
	  "species_id": 226,
	  "height": 21,
	  "weight": 2200,
	  "base_experience": 170,
	  "order": 305,
	  "is_default": true
	},
	{
	  "id": 227,
	  "identifier": "skarmory",
	  "species_id": 227,
	  "height": 17,
	  "weight": 505,
	  "base_experience": 163,
	  "order": 306,
	  "is_default": true
	},
	{
	  "id": 228,
	  "identifier": "houndour",
	  "species_id": 228,
	  "height": 6,
	  "weight": 108,
	  "base_experience": 66,
	  "order": 307,
	  "is_default": true
	},
	{
	  "id": 229,
	  "identifier": "houndoom",
	  "species_id": 229,
	  "height": 14,
	  "weight": 350,
	  "base_experience": 175,
	  "order": 308,
	  "is_default": true
	},
	{
	  "id": 230,
	  "identifier": "kingdra",
	  "species_id": 230,
	  "height": 18,
	  "weight": 1520,
	  "base_experience": 243,
	  "order": 172,
	  "is_default": true
	},
	{
	  "id": 231,
	  "identifier": "phanpy",
	  "species_id": 231,
	  "height": 5,
	  "weight": 335,
	  "base_experience": 66,
	  "order": 310,
	  "is_default": true
	},
	{
	  "id": 232,
	  "identifier": "donphan",
	  "species_id": 232,
	  "height": 11,
	  "weight": 1200,
	  "base_experience": 175,
	  "order": 311,
	  "is_default": true
	},
	{
	  "id": 233,
	  "identifier": "porygon2",
	  "species_id": 233,
	  "height": 6,
	  "weight": 325,
	  "base_experience": 180,
	  "order": 208,
	  "is_default": true
	},
	{
	  "id": 234,
	  "identifier": "stantler",
	  "species_id": 234,
	  "height": 14,
	  "weight": 712,
	  "base_experience": 163,
	  "order": 312,
	  "is_default": true
	},
	{
	  "id": 235,
	  "identifier": "smeargle",
	  "species_id": 235,
	  "height": 12,
	  "weight": 580,
	  "base_experience": 88,
	  "order": 313,
	  "is_default": true
	},
	{
	  "id": 236,
	  "identifier": "tyrogue",
	  "species_id": 236,
	  "height": 7,
	  "weight": 210,
	  "base_experience": 42,
	  "order": 152,
	  "is_default": true
	},
	{
	  "id": 237,
	  "identifier": "hitmontop",
	  "species_id": 237,
	  "height": 14,
	  "weight": 480,
	  "base_experience": 159,
	  "order": 155,
	  "is_default": true
	},
	{
	  "id": 238,
	  "identifier": "smoochum",
	  "species_id": 238,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 61,
	  "order": 182,
	  "is_default": true
	},
	{
	  "id": 239,
	  "identifier": "elekid",
	  "species_id": 239,
	  "height": 6,
	  "weight": 235,
	  "base_experience": 72,
	  "order": 184,
	  "is_default": true
	},
	{
	  "id": 240,
	  "identifier": "magby",
	  "species_id": 240,
	  "height": 7,
	  "weight": 214,
	  "base_experience": 73,
	  "order": 187,
	  "is_default": true
	},
	{
	  "id": 241,
	  "identifier": "miltank",
	  "species_id": 241,
	  "height": 12,
	  "weight": 755,
	  "base_experience": 172,
	  "order": 314,
	  "is_default": true
	},
	{
	  "id": 242,
	  "identifier": "blissey",
	  "species_id": 242,
	  "height": 15,
	  "weight": 468,
	  "base_experience": 608,
	  "order": 165,
	  "is_default": true
	},
	{
	  "id": 243,
	  "identifier": "raikou",
	  "species_id": 243,
	  "height": 19,
	  "weight": 1780,
	  "base_experience": 261,
	  "order": 315,
	  "is_default": true
	},
	{
	  "id": 244,
	  "identifier": "entei",
	  "species_id": 244,
	  "height": 21,
	  "weight": 1980,
	  "base_experience": 261,
	  "order": 316,
	  "is_default": true
	},
	{
	  "id": 245,
	  "identifier": "suicune",
	  "species_id": 245,
	  "height": 20,
	  "weight": 1870,
	  "base_experience": 261,
	  "order": 317,
	  "is_default": true
	},
	{
	  "id": 246,
	  "identifier": "larvitar",
	  "species_id": 246,
	  "height": 6,
	  "weight": 720,
	  "base_experience": 60,
	  "order": 318,
	  "is_default": true
	},
	{
	  "id": 247,
	  "identifier": "pupitar",
	  "species_id": 247,
	  "height": 12,
	  "weight": 1520,
	  "base_experience": 144,
	  "order": 319,
	  "is_default": true
	},
	{
	  "id": 248,
	  "identifier": "tyranitar",
	  "species_id": 248,
	  "height": 20,
	  "weight": 2020,
	  "base_experience": 270,
	  "order": 320,
	  "is_default": true
	},
	{
	  "id": 249,
	  "identifier": "lugia",
	  "species_id": 249,
	  "height": 52,
	  "weight": 2160,
	  "base_experience": 306,
	  "order": 322,
	  "is_default": true
	},
	{
	  "id": 250,
	  "identifier": "ho-oh",
	  "species_id": 250,
	  "height": 38,
	  "weight": 1990,
	  "base_experience": 306,
	  "order": 323,
	  "is_default": true
	},
	{
	  "id": 251,
	  "identifier": "celebi",
	  "species_id": 251,
	  "height": 6,
	  "weight": 50,
	  "base_experience": 270,
	  "order": 324,
	  "is_default": true
	},
	{
	  "id": 252,
	  "identifier": "treecko",
	  "species_id": 252,
	  "height": 5,
	  "weight": 50,
	  "base_experience": 62,
	  "order": 325,
	  "is_default": true
	},
	{
	  "id": 253,
	  "identifier": "grovyle",
	  "species_id": 253,
	  "height": 9,
	  "weight": 216,
	  "base_experience": 142,
	  "order": 326,
	  "is_default": true
	},
	{
	  "id": 254,
	  "identifier": "sceptile",
	  "species_id": 254,
	  "height": 17,
	  "weight": 522,
	  "base_experience": 239,
	  "order": 327,
	  "is_default": true
	},
	{
	  "id": 255,
	  "identifier": "torchic",
	  "species_id": 255,
	  "height": 4,
	  "weight": 25,
	  "base_experience": 62,
	  "order": 329,
	  "is_default": true
	},
	{
	  "id": 256,
	  "identifier": "combusken",
	  "species_id": 256,
	  "height": 9,
	  "weight": 195,
	  "base_experience": 142,
	  "order": 330,
	  "is_default": true
	},
	{
	  "id": 257,
	  "identifier": "blaziken",
	  "species_id": 257,
	  "height": 19,
	  "weight": 520,
	  "base_experience": 239,
	  "order": 331,
	  "is_default": true
	},
	{
	  "id": 258,
	  "identifier": "mudkip",
	  "species_id": 258,
	  "height": 4,
	  "weight": 76,
	  "base_experience": 62,
	  "order": 333,
	  "is_default": true
	},
	{
	  "id": 259,
	  "identifier": "marshtomp",
	  "species_id": 259,
	  "height": 7,
	  "weight": 280,
	  "base_experience": 142,
	  "order": 334,
	  "is_default": true
	},
	{
	  "id": 260,
	  "identifier": "swampert",
	  "species_id": 260,
	  "height": 15,
	  "weight": 819,
	  "base_experience": 241,
	  "order": 335,
	  "is_default": true
	},
	{
	  "id": 261,
	  "identifier": "poochyena",
	  "species_id": 261,
	  "height": 5,
	  "weight": 136,
	  "base_experience": 56,
	  "order": 337,
	  "is_default": true
	},
	{
	  "id": 262,
	  "identifier": "mightyena",
	  "species_id": 262,
	  "height": 10,
	  "weight": 370,
	  "base_experience": 147,
	  "order": 338,
	  "is_default": true
	},
	{
	  "id": 263,
	  "identifier": "zigzagoon",
	  "species_id": 263,
	  "height": 4,
	  "weight": 175,
	  "base_experience": 56,
	  "order": 339,
	  "is_default": true
	},
	{
	  "id": 264,
	  "identifier": "linoone",
	  "species_id": 264,
	  "height": 5,
	  "weight": 325,
	  "base_experience": 147,
	  "order": 340,
	  "is_default": true
	},
	{
	  "id": 265,
	  "identifier": "wurmple",
	  "species_id": 265,
	  "height": 3,
	  "weight": 36,
	  "base_experience": 56,
	  "order": 341,
	  "is_default": true
	},
	{
	  "id": 266,
	  "identifier": "silcoon",
	  "species_id": 266,
	  "height": 6,
	  "weight": 100,
	  "base_experience": 72,
	  "order": 342,
	  "is_default": true
	},
	{
	  "id": 267,
	  "identifier": "beautifly",
	  "species_id": 267,
	  "height": 10,
	  "weight": 284,
	  "base_experience": 178,
	  "order": 343,
	  "is_default": true
	},
	{
	  "id": 268,
	  "identifier": "cascoon",
	  "species_id": 268,
	  "height": 7,
	  "weight": 115,
	  "base_experience": 72,
	  "order": 344,
	  "is_default": true
	},
	{
	  "id": 269,
	  "identifier": "dustox",
	  "species_id": 269,
	  "height": 12,
	  "weight": 316,
	  "base_experience": 173,
	  "order": 345,
	  "is_default": true
	},
	{
	  "id": 270,
	  "identifier": "lotad",
	  "species_id": 270,
	  "height": 5,
	  "weight": 26,
	  "base_experience": 44,
	  "order": 346,
	  "is_default": true
	},
	{
	  "id": 271,
	  "identifier": "lombre",
	  "species_id": 271,
	  "height": 12,
	  "weight": 325,
	  "base_experience": 119,
	  "order": 347,
	  "is_default": true
	},
	{
	  "id": 272,
	  "identifier": "ludicolo",
	  "species_id": 272,
	  "height": 15,
	  "weight": 550,
	  "base_experience": 216,
	  "order": 348,
	  "is_default": true
	},
	{
	  "id": 273,
	  "identifier": "seedot",
	  "species_id": 273,
	  "height": 5,
	  "weight": 40,
	  "base_experience": 44,
	  "order": 349,
	  "is_default": true
	},
	{
	  "id": 274,
	  "identifier": "nuzleaf",
	  "species_id": 274,
	  "height": 10,
	  "weight": 280,
	  "base_experience": 119,
	  "order": 350,
	  "is_default": true
	},
	{
	  "id": 275,
	  "identifier": "shiftry",
	  "species_id": 275,
	  "height": 13,
	  "weight": 596,
	  "base_experience": 216,
	  "order": 351,
	  "is_default": true
	},
	{
	  "id": 276,
	  "identifier": "taillow",
	  "species_id": 276,
	  "height": 3,
	  "weight": 23,
	  "base_experience": 54,
	  "order": 352,
	  "is_default": true
	},
	{
	  "id": 277,
	  "identifier": "swellow",
	  "species_id": 277,
	  "height": 7,
	  "weight": 198,
	  "base_experience": 159,
	  "order": 353,
	  "is_default": true
	},
	{
	  "id": 278,
	  "identifier": "wingull",
	  "species_id": 278,
	  "height": 6,
	  "weight": 95,
	  "base_experience": 54,
	  "order": 354,
	  "is_default": true
	},
	{
	  "id": 279,
	  "identifier": "pelipper",
	  "species_id": 279,
	  "height": 12,
	  "weight": 280,
	  "base_experience": 154,
	  "order": 355,
	  "is_default": true
	},
	{
	  "id": 280,
	  "identifier": "ralts",
	  "species_id": 280,
	  "height": 4,
	  "weight": 66,
	  "base_experience": 40,
	  "order": 356,
	  "is_default": true
	},
	{
	  "id": 281,
	  "identifier": "kirlia",
	  "species_id": 281,
	  "height": 8,
	  "weight": 202,
	  "base_experience": 97,
	  "order": 357,
	  "is_default": true
	},
	{
	  "id": 282,
	  "identifier": "gardevoir",
	  "species_id": 282,
	  "height": 16,
	  "weight": 484,
	  "base_experience": 233,
	  "order": 358,
	  "is_default": true
	},
	{
	  "id": 283,
	  "identifier": "surskit",
	  "species_id": 283,
	  "height": 5,
	  "weight": 17,
	  "base_experience": 54,
	  "order": 362,
	  "is_default": true
	},
	{
	  "id": 284,
	  "identifier": "masquerain",
	  "species_id": 284,
	  "height": 8,
	  "weight": 36,
	  "base_experience": 159,
	  "order": 363,
	  "is_default": true
	},
	{
	  "id": 285,
	  "identifier": "shroomish",
	  "species_id": 285,
	  "height": 4,
	  "weight": 45,
	  "base_experience": 59,
	  "order": 364,
	  "is_default": true
	},
	{
	  "id": 286,
	  "identifier": "breloom",
	  "species_id": 286,
	  "height": 12,
	  "weight": 392,
	  "base_experience": 161,
	  "order": 365,
	  "is_default": true
	},
	{
	  "id": 287,
	  "identifier": "slakoth",
	  "species_id": 287,
	  "height": 8,
	  "weight": 240,
	  "base_experience": 56,
	  "order": 366,
	  "is_default": true
	},
	{
	  "id": 288,
	  "identifier": "vigoroth",
	  "species_id": 288,
	  "height": 14,
	  "weight": 465,
	  "base_experience": 154,
	  "order": 367,
	  "is_default": true
	},
	{
	  "id": 289,
	  "identifier": "slaking",
	  "species_id": 289,
	  "height": 20,
	  "weight": 1305,
	  "base_experience": 252,
	  "order": 368,
	  "is_default": true
	},
	{
	  "id": 290,
	  "identifier": "nincada",
	  "species_id": 290,
	  "height": 5,
	  "weight": 55,
	  "base_experience": 53,
	  "order": 369,
	  "is_default": true
	},
	{
	  "id": 291,
	  "identifier": "ninjask",
	  "species_id": 291,
	  "height": 8,
	  "weight": 120,
	  "base_experience": 160,
	  "order": 370,
	  "is_default": true
	},
	{
	  "id": 292,
	  "identifier": "shedinja",
	  "species_id": 292,
	  "height": 8,
	  "weight": 12,
	  "base_experience": 83,
	  "order": 371,
	  "is_default": true
	},
	{
	  "id": 293,
	  "identifier": "whismur",
	  "species_id": 293,
	  "height": 6,
	  "weight": 163,
	  "base_experience": 48,
	  "order": 372,
	  "is_default": true
	},
	{
	  "id": 294,
	  "identifier": "loudred",
	  "species_id": 294,
	  "height": 10,
	  "weight": 405,
	  "base_experience": 126,
	  "order": 373,
	  "is_default": true
	},
	{
	  "id": 295,
	  "identifier": "exploud",
	  "species_id": 295,
	  "height": 15,
	  "weight": 840,
	  "base_experience": 221,
	  "order": 374,
	  "is_default": true
	},
	{
	  "id": 296,
	  "identifier": "makuhita",
	  "species_id": 296,
	  "height": 10,
	  "weight": 864,
	  "base_experience": 47,
	  "order": 375,
	  "is_default": true
	},
	{
	  "id": 297,
	  "identifier": "hariyama",
	  "species_id": 297,
	  "height": 23,
	  "weight": 2538,
	  "base_experience": 166,
	  "order": 376,
	  "is_default": true
	},
	{
	  "id": 298,
	  "identifier": "azurill",
	  "species_id": 298,
	  "height": 2,
	  "weight": 20,
	  "base_experience": 38,
	  "order": 256,
	  "is_default": true
	},
	{
	  "id": 299,
	  "identifier": "nosepass",
	  "species_id": 299,
	  "height": 10,
	  "weight": 970,
	  "base_experience": 75,
	  "order": 377,
	  "is_default": true
	},
	{
	  "id": 300,
	  "identifier": "skitty",
	  "species_id": 300,
	  "height": 6,
	  "weight": 110,
	  "base_experience": 52,
	  "order": 379,
	  "is_default": true
	},
	{
	  "id": 301,
	  "identifier": "delcatty",
	  "species_id": 301,
	  "height": 11,
	  "weight": 326,
	  "base_experience": 140,
	  "order": 380,
	  "is_default": true
	},
	{
	  "id": 302,
	  "identifier": "sableye",
	  "species_id": 302,
	  "height": 5,
	  "weight": 110,
	  "base_experience": 133,
	  "order": 381,
	  "is_default": true
	},
	{
	  "id": 303,
	  "identifier": "mawile",
	  "species_id": 303,
	  "height": 6,
	  "weight": 115,
	  "base_experience": 133,
	  "order": 383,
	  "is_default": true
	},
	{
	  "id": 304,
	  "identifier": "aron",
	  "species_id": 304,
	  "height": 4,
	  "weight": 600,
	  "base_experience": 66,
	  "order": 385,
	  "is_default": true
	},
	{
	  "id": 305,
	  "identifier": "lairon",
	  "species_id": 305,
	  "height": 9,
	  "weight": 1200,
	  "base_experience": 151,
	  "order": 386,
	  "is_default": true
	},
	{
	  "id": 306,
	  "identifier": "aggron",
	  "species_id": 306,
	  "height": 21,
	  "weight": 3600,
	  "base_experience": 239,
	  "order": 387,
	  "is_default": true
	},
	{
	  "id": 307,
	  "identifier": "meditite",
	  "species_id": 307,
	  "height": 6,
	  "weight": 112,
	  "base_experience": 56,
	  "order": 389,
	  "is_default": true
	},
	{
	  "id": 308,
	  "identifier": "medicham",
	  "species_id": 308,
	  "height": 13,
	  "weight": 315,
	  "base_experience": 144,
	  "order": 390,
	  "is_default": true
	},
	{
	  "id": 309,
	  "identifier": "electrike",
	  "species_id": 309,
	  "height": 6,
	  "weight": 152,
	  "base_experience": 59,
	  "order": 392,
	  "is_default": true
	},
	{
	  "id": 310,
	  "identifier": "manectric",
	  "species_id": 310,
	  "height": 15,
	  "weight": 402,
	  "base_experience": 166,
	  "order": 393,
	  "is_default": true
	},
	{
	  "id": 311,
	  "identifier": "plusle",
	  "species_id": 311,
	  "height": 4,
	  "weight": 42,
	  "base_experience": 142,
	  "order": 395,
	  "is_default": true
	},
	{
	  "id": 312,
	  "identifier": "minun",
	  "species_id": 312,
	  "height": 4,
	  "weight": 42,
	  "base_experience": 142,
	  "order": 396,
	  "is_default": true
	},
	{
	  "id": 313,
	  "identifier": "volbeat",
	  "species_id": 313,
	  "height": 7,
	  "weight": 177,
	  "base_experience": 151,
	  "order": 397,
	  "is_default": true
	},
	{
	  "id": 314,
	  "identifier": "illumise",
	  "species_id": 314,
	  "height": 6,
	  "weight": 177,
	  "base_experience": 151,
	  "order": 398,
	  "is_default": true
	},
	{
	  "id": 315,
	  "identifier": "roselia",
	  "species_id": 315,
	  "height": 3,
	  "weight": 20,
	  "base_experience": 140,
	  "order": 400,
	  "is_default": true
	},
	{
	  "id": 316,
	  "identifier": "gulpin",
	  "species_id": 316,
	  "height": 4,
	  "weight": 103,
	  "base_experience": 60,
	  "order": 402,
	  "is_default": true
	},
	{
	  "id": 317,
	  "identifier": "swalot",
	  "species_id": 317,
	  "height": 17,
	  "weight": 800,
	  "base_experience": 163,
	  "order": 403,
	  "is_default": true
	},
	{
	  "id": 318,
	  "identifier": "carvanha",
	  "species_id": 318,
	  "height": 8,
	  "weight": 208,
	  "base_experience": 61,
	  "order": 404,
	  "is_default": true
	},
	{
	  "id": 319,
	  "identifier": "sharpedo",
	  "species_id": 319,
	  "height": 18,
	  "weight": 888,
	  "base_experience": 161,
	  "order": 405,
	  "is_default": true
	},
	{
	  "id": 320,
	  "identifier": "wailmer",
	  "species_id": 320,
	  "height": 20,
	  "weight": 1300,
	  "base_experience": 80,
	  "order": 407,
	  "is_default": true
	},
	{
	  "id": 321,
	  "identifier": "wailord",
	  "species_id": 321,
	  "height": 145,
	  "weight": 3980,
	  "base_experience": 175,
	  "order": 408,
	  "is_default": true
	},
	{
	  "id": 322,
	  "identifier": "numel",
	  "species_id": 322,
	  "height": 7,
	  "weight": 240,
	  "base_experience": 61,
	  "order": 409,
	  "is_default": true
	},
	{
	  "id": 323,
	  "identifier": "camerupt",
	  "species_id": 323,
	  "height": 19,
	  "weight": 2200,
	  "base_experience": 161,
	  "order": 410,
	  "is_default": true
	},
	{
	  "id": 324,
	  "identifier": "torkoal",
	  "species_id": 324,
	  "height": 5,
	  "weight": 804,
	  "base_experience": 165,
	  "order": 412,
	  "is_default": true
	},
	{
	  "id": 325,
	  "identifier": "spoink",
	  "species_id": 325,
	  "height": 7,
	  "weight": 306,
	  "base_experience": 66,
	  "order": 413,
	  "is_default": true
	},
	{
	  "id": 326,
	  "identifier": "grumpig",
	  "species_id": 326,
	  "height": 9,
	  "weight": 715,
	  "base_experience": 165,
	  "order": 414,
	  "is_default": true
	},
	{
	  "id": 327,
	  "identifier": "spinda",
	  "species_id": 327,
	  "height": 11,
	  "weight": 50,
	  "base_experience": 126,
	  "order": 415,
	  "is_default": true
	},
	{
	  "id": 328,
	  "identifier": "trapinch",
	  "species_id": 328,
	  "height": 7,
	  "weight": 150,
	  "base_experience": 58,
	  "order": 416,
	  "is_default": true
	},
	{
	  "id": 329,
	  "identifier": "vibrava",
	  "species_id": 329,
	  "height": 11,
	  "weight": 153,
	  "base_experience": 119,
	  "order": 417,
	  "is_default": true
	},
	{
	  "id": 330,
	  "identifier": "flygon",
	  "species_id": 330,
	  "height": 20,
	  "weight": 820,
	  "base_experience": 234,
	  "order": 418,
	  "is_default": true
	},
	{
	  "id": 331,
	  "identifier": "cacnea",
	  "species_id": 331,
	  "height": 4,
	  "weight": 513,
	  "base_experience": 67,
	  "order": 419,
	  "is_default": true
	},
	{
	  "id": 332,
	  "identifier": "cacturne",
	  "species_id": 332,
	  "height": 13,
	  "weight": 774,
	  "base_experience": 166,
	  "order": 420,
	  "is_default": true
	},
	{
	  "id": 333,
	  "identifier": "swablu",
	  "species_id": 333,
	  "height": 4,
	  "weight": 12,
	  "base_experience": 62,
	  "order": 421,
	  "is_default": true
	},
	{
	  "id": 334,
	  "identifier": "altaria",
	  "species_id": 334,
	  "height": 11,
	  "weight": 206,
	  "base_experience": 172,
	  "order": 422,
	  "is_default": true
	},
	{
	  "id": 335,
	  "identifier": "zangoose",
	  "species_id": 335,
	  "height": 13,
	  "weight": 403,
	  "base_experience": 160,
	  "order": 424,
	  "is_default": true
	},
	{
	  "id": 336,
	  "identifier": "seviper",
	  "species_id": 336,
	  "height": 27,
	  "weight": 525,
	  "base_experience": 160,
	  "order": 425,
	  "is_default": true
	},
	{
	  "id": 337,
	  "identifier": "lunatone",
	  "species_id": 337,
	  "height": 10,
	  "weight": 1680,
	  "base_experience": 161,
	  "order": 426,
	  "is_default": true
	},
	{
	  "id": 338,
	  "identifier": "solrock",
	  "species_id": 338,
	  "height": 12,
	  "weight": 1540,
	  "base_experience": 161,
	  "order": 427,
	  "is_default": true
	},
	{
	  "id": 339,
	  "identifier": "barboach",
	  "species_id": 339,
	  "height": 4,
	  "weight": 19,
	  "base_experience": 58,
	  "order": 428,
	  "is_default": true
	},
	{
	  "id": 340,
	  "identifier": "whiscash",
	  "species_id": 340,
	  "height": 9,
	  "weight": 236,
	  "base_experience": 164,
	  "order": 429,
	  "is_default": true
	},
	{
	  "id": 341,
	  "identifier": "corphish",
	  "species_id": 341,
	  "height": 6,
	  "weight": 115,
	  "base_experience": 62,
	  "order": 430,
	  "is_default": true
	},
	{
	  "id": 342,
	  "identifier": "crawdaunt",
	  "species_id": 342,
	  "height": 11,
	  "weight": 328,
	  "base_experience": 164,
	  "order": 431,
	  "is_default": true
	},
	{
	  "id": 343,
	  "identifier": "baltoy",
	  "species_id": 343,
	  "height": 5,
	  "weight": 215,
	  "base_experience": 60,
	  "order": 432,
	  "is_default": true
	},
	{
	  "id": 344,
	  "identifier": "claydol",
	  "species_id": 344,
	  "height": 15,
	  "weight": 1080,
	  "base_experience": 175,
	  "order": 433,
	  "is_default": true
	},
	{
	  "id": 345,
	  "identifier": "lileep",
	  "species_id": 345,
	  "height": 10,
	  "weight": 238,
	  "base_experience": 71,
	  "order": 434,
	  "is_default": true
	},
	{
	  "id": 346,
	  "identifier": "cradily",
	  "species_id": 346,
	  "height": 15,
	  "weight": 604,
	  "base_experience": 173,
	  "order": 435,
	  "is_default": true
	},
	{
	  "id": 347,
	  "identifier": "anorith",
	  "species_id": 347,
	  "height": 7,
	  "weight": 125,
	  "base_experience": 71,
	  "order": 436,
	  "is_default": true
	},
	{
	  "id": 348,
	  "identifier": "armaldo",
	  "species_id": 348,
	  "height": 15,
	  "weight": 682,
	  "base_experience": 173,
	  "order": 437,
	  "is_default": true
	},
	{
	  "id": 349,
	  "identifier": "feebas",
	  "species_id": 349,
	  "height": 6,
	  "weight": 74,
	  "base_experience": 40,
	  "order": 438,
	  "is_default": true
	},
	{
	  "id": 350,
	  "identifier": "milotic",
	  "species_id": 350,
	  "height": 62,
	  "weight": 1620,
	  "base_experience": 189,
	  "order": 439,
	  "is_default": true
	},
	{
	  "id": 351,
	  "identifier": "castform",
	  "species_id": 351,
	  "height": 3,
	  "weight": 8,
	  "base_experience": 147,
	  "order": 440,
	  "is_default": true
	},
	{
	  "id": 352,
	  "identifier": "kecleon",
	  "species_id": 352,
	  "height": 10,
	  "weight": 220,
	  "base_experience": 154,
	  "order": 444,
	  "is_default": true
	},
	{
	  "id": 353,
	  "identifier": "shuppet",
	  "species_id": 353,
	  "height": 6,
	  "weight": 23,
	  "base_experience": 59,
	  "order": 445,
	  "is_default": true
	},
	{
	  "id": 354,
	  "identifier": "banette",
	  "species_id": 354,
	  "height": 11,
	  "weight": 125,
	  "base_experience": 159,
	  "order": 446,
	  "is_default": true
	},
	{
	  "id": 355,
	  "identifier": "duskull",
	  "species_id": 355,
	  "height": 8,
	  "weight": 150,
	  "base_experience": 59,
	  "order": 448,
	  "is_default": true
	},
	{
	  "id": 356,
	  "identifier": "dusclops",
	  "species_id": 356,
	  "height": 16,
	  "weight": 306,
	  "base_experience": 159,
	  "order": 449,
	  "is_default": true
	},
	{
	  "id": 357,
	  "identifier": "tropius",
	  "species_id": 357,
	  "height": 20,
	  "weight": 1000,
	  "base_experience": 161,
	  "order": 451,
	  "is_default": true
	},
	{
	  "id": 358,
	  "identifier": "chimecho",
	  "species_id": 358,
	  "height": 6,
	  "weight": 10,
	  "base_experience": 159,
	  "order": 453,
	  "is_default": true
	},
	{
	  "id": 359,
	  "identifier": "absol",
	  "species_id": 359,
	  "height": 12,
	  "weight": 470,
	  "base_experience": 163,
	  "order": 454,
	  "is_default": true
	},
	{
	  "id": 360,
	  "identifier": "wynaut",
	  "species_id": 360,
	  "height": 6,
	  "weight": 140,
	  "base_experience": 52,
	  "order": 277,
	  "is_default": true
	},
	{
	  "id": 361,
	  "identifier": "snorunt",
	  "species_id": 361,
	  "height": 7,
	  "weight": 168,
	  "base_experience": 60,
	  "order": 456,
	  "is_default": true
	},
	{
	  "id": 362,
	  "identifier": "glalie",
	  "species_id": 362,
	  "height": 15,
	  "weight": 2565,
	  "base_experience": 168,
	  "order": 457,
	  "is_default": true
	},
	{
	  "id": 363,
	  "identifier": "spheal",
	  "species_id": 363,
	  "height": 8,
	  "weight": 395,
	  "base_experience": 58,
	  "order": 460,
	  "is_default": true
	},
	{
	  "id": 364,
	  "identifier": "sealeo",
	  "species_id": 364,
	  "height": 11,
	  "weight": 876,
	  "base_experience": 144,
	  "order": 461,
	  "is_default": true
	},
	{
	  "id": 365,
	  "identifier": "walrein",
	  "species_id": 365,
	  "height": 14,
	  "weight": 1506,
	  "base_experience": 239,
	  "order": 462,
	  "is_default": true
	},
	{
	  "id": 366,
	  "identifier": "clamperl",
	  "species_id": 366,
	  "height": 4,
	  "weight": 525,
	  "base_experience": 69,
	  "order": 463,
	  "is_default": true
	},
	{
	  "id": 367,
	  "identifier": "huntail",
	  "species_id": 367,
	  "height": 17,
	  "weight": 270,
	  "base_experience": 170,
	  "order": 464,
	  "is_default": true
	},
	{
	  "id": 368,
	  "identifier": "gorebyss",
	  "species_id": 368,
	  "height": 18,
	  "weight": 226,
	  "base_experience": 170,
	  "order": 465,
	  "is_default": true
	},
	{
	  "id": 369,
	  "identifier": "relicanth",
	  "species_id": 369,
	  "height": 10,
	  "weight": 234,
	  "base_experience": 170,
	  "order": 466,
	  "is_default": true
	},
	{
	  "id": 370,
	  "identifier": "luvdisc",
	  "species_id": 370,
	  "height": 6,
	  "weight": 87,
	  "base_experience": 116,
	  "order": 467,
	  "is_default": true
	},
	{
	  "id": 371,
	  "identifier": "bagon",
	  "species_id": 371,
	  "height": 6,
	  "weight": 421,
	  "base_experience": 60,
	  "order": 468,
	  "is_default": true
	},
	{
	  "id": 372,
	  "identifier": "shelgon",
	  "species_id": 372,
	  "height": 11,
	  "weight": 1105,
	  "base_experience": 147,
	  "order": 469,
	  "is_default": true
	},
	{
	  "id": 373,
	  "identifier": "salamence",
	  "species_id": 373,
	  "height": 15,
	  "weight": 1026,
	  "base_experience": 270,
	  "order": 470,
	  "is_default": true
	},
	{
	  "id": 374,
	  "identifier": "beldum",
	  "species_id": 374,
	  "height": 6,
	  "weight": 952,
	  "base_experience": 60,
	  "order": 472,
	  "is_default": true
	},
	{
	  "id": 375,
	  "identifier": "metang",
	  "species_id": 375,
	  "height": 12,
	  "weight": 2025,
	  "base_experience": 147,
	  "order": 473,
	  "is_default": true
	},
	{
	  "id": 376,
	  "identifier": "metagross",
	  "species_id": 376,
	  "height": 16,
	  "weight": 5500,
	  "base_experience": 270,
	  "order": 474,
	  "is_default": true
	},
	{
	  "id": 377,
	  "identifier": "regirock",
	  "species_id": 377,
	  "height": 17,
	  "weight": 2300,
	  "base_experience": 261,
	  "order": 476,
	  "is_default": true
	},
	{
	  "id": 378,
	  "identifier": "regice",
	  "species_id": 378,
	  "height": 18,
	  "weight": 1750,
	  "base_experience": 261,
	  "order": 477,
	  "is_default": true
	},
	{
	  "id": 379,
	  "identifier": "registeel",
	  "species_id": 379,
	  "height": 19,
	  "weight": 2050,
	  "base_experience": 261,
	  "order": 478,
	  "is_default": true
	},
	{
	  "id": 380,
	  "identifier": "latias",
	  "species_id": 380,
	  "height": 14,
	  "weight": 400,
	  "base_experience": 270,
	  "order": 479,
	  "is_default": true
	},
	{
	  "id": 381,
	  "identifier": "latios",
	  "species_id": 381,
	  "height": 20,
	  "weight": 600,
	  "base_experience": 270,
	  "order": 481,
	  "is_default": true
	},
	{
	  "id": 382,
	  "identifier": "kyogre",
	  "species_id": 382,
	  "height": 45,
	  "weight": 3520,
	  "base_experience": 302,
	  "order": 483,
	  "is_default": true
	},
	{
	  "id": 383,
	  "identifier": "groudon",
	  "species_id": 383,
	  "height": 35,
	  "weight": 9500,
	  "base_experience": 302,
	  "order": 485,
	  "is_default": true
	},
	{
	  "id": 384,
	  "identifier": "rayquaza",
	  "species_id": 384,
	  "height": 70,
	  "weight": 2065,
	  "base_experience": 306,
	  "order": 487,
	  "is_default": true
	},
	{
	  "id": 385,
	  "identifier": "jirachi",
	  "species_id": 385,
	  "height": 3,
	  "weight": 11,
	  "base_experience": 270,
	  "order": 489,
	  "is_default": true
	},
	{
	  "id": 386,
	  "identifier": "deoxys-normal",
	  "species_id": 386,
	  "height": 17,
	  "weight": 608,
	  "base_experience": 270,
	  "order": 490,
	  "is_default": true
	},
	{
	  "id": 387,
	  "identifier": "turtwig",
	  "species_id": 387,
	  "height": 4,
	  "weight": 102,
	  "base_experience": 64,
	  "order": 494,
	  "is_default": true
	},
	{
	  "id": 388,
	  "identifier": "grotle",
	  "species_id": 388,
	  "height": 11,
	  "weight": 970,
	  "base_experience": 142,
	  "order": 495,
	  "is_default": true
	},
	{
	  "id": 389,
	  "identifier": "torterra",
	  "species_id": 389,
	  "height": 22,
	  "weight": 3100,
	  "base_experience": 236,
	  "order": 496,
	  "is_default": true
	},
	{
	  "id": 390,
	  "identifier": "chimchar",
	  "species_id": 390,
	  "height": 5,
	  "weight": 62,
	  "base_experience": 62,
	  "order": 497,
	  "is_default": true
	},
	{
	  "id": 391,
	  "identifier": "monferno",
	  "species_id": 391,
	  "height": 9,
	  "weight": 220,
	  "base_experience": 142,
	  "order": 498,
	  "is_default": true
	},
	{
	  "id": 392,
	  "identifier": "infernape",
	  "species_id": 392,
	  "height": 12,
	  "weight": 550,
	  "base_experience": 240,
	  "order": 499,
	  "is_default": true
	},
	{
	  "id": 393,
	  "identifier": "piplup",
	  "species_id": 393,
	  "height": 4,
	  "weight": 52,
	  "base_experience": 63,
	  "order": 500,
	  "is_default": true
	},
	{
	  "id": 394,
	  "identifier": "prinplup",
	  "species_id": 394,
	  "height": 8,
	  "weight": 230,
	  "base_experience": 142,
	  "order": 501,
	  "is_default": true
	},
	{
	  "id": 395,
	  "identifier": "empoleon",
	  "species_id": 395,
	  "height": 17,
	  "weight": 845,
	  "base_experience": 239,
	  "order": 502,
	  "is_default": true
	},
	{
	  "id": 396,
	  "identifier": "starly",
	  "species_id": 396,
	  "height": 3,
	  "weight": 20,
	  "base_experience": 49,
	  "order": 503,
	  "is_default": true
	},
	{
	  "id": 397,
	  "identifier": "staravia",
	  "species_id": 397,
	  "height": 6,
	  "weight": 155,
	  "base_experience": 119,
	  "order": 504,
	  "is_default": true
	},
	{
	  "id": 398,
	  "identifier": "staraptor",
	  "species_id": 398,
	  "height": 12,
	  "weight": 249,
	  "base_experience": 218,
	  "order": 505,
	  "is_default": true
	},
	{
	  "id": 399,
	  "identifier": "bidoof",
	  "species_id": 399,
	  "height": 5,
	  "weight": 200,
	  "base_experience": 50,
	  "order": 506,
	  "is_default": true
	},
	{
	  "id": 400,
	  "identifier": "bibarel",
	  "species_id": 400,
	  "height": 10,
	  "weight": 315,
	  "base_experience": 144,
	  "order": 507,
	  "is_default": true
	},
	{
	  "id": 401,
	  "identifier": "kricketot",
	  "species_id": 401,
	  "height": 3,
	  "weight": 22,
	  "base_experience": 39,
	  "order": 508,
	  "is_default": true
	},
	{
	  "id": 402,
	  "identifier": "kricketune",
	  "species_id": 402,
	  "height": 10,
	  "weight": 255,
	  "base_experience": 134,
	  "order": 509,
	  "is_default": true
	},
	{
	  "id": 403,
	  "identifier": "shinx",
	  "species_id": 403,
	  "height": 5,
	  "weight": 95,
	  "base_experience": 53,
	  "order": 510,
	  "is_default": true
	},
	{
	  "id": 404,
	  "identifier": "luxio",
	  "species_id": 404,
	  "height": 9,
	  "weight": 305,
	  "base_experience": 127,
	  "order": 511,
	  "is_default": true
	},
	{
	  "id": 405,
	  "identifier": "luxray",
	  "species_id": 405,
	  "height": 14,
	  "weight": 420,
	  "base_experience": 235,
	  "order": 512,
	  "is_default": true
	},
	{
	  "id": 406,
	  "identifier": "budew",
	  "species_id": 406,
	  "height": 2,
	  "weight": 12,
	  "base_experience": 56,
	  "order": 399,
	  "is_default": true
	},
	{
	  "id": 407,
	  "identifier": "roserade",
	  "species_id": 407,
	  "height": 9,
	  "weight": 145,
	  "base_experience": 232,
	  "order": 401,
	  "is_default": true
	},
	{
	  "id": 408,
	  "identifier": "cranidos",
	  "species_id": 408,
	  "height": 9,
	  "weight": 315,
	  "base_experience": 70,
	  "order": 513,
	  "is_default": true
	},
	{
	  "id": 409,
	  "identifier": "rampardos",
	  "species_id": 409,
	  "height": 16,
	  "weight": 1025,
	  "base_experience": 173,
	  "order": 514,
	  "is_default": true
	},
	{
	  "id": 410,
	  "identifier": "shieldon",
	  "species_id": 410,
	  "height": 5,
	  "weight": 570,
	  "base_experience": 70,
	  "order": 515,
	  "is_default": true
	},
	{
	  "id": 411,
	  "identifier": "bastiodon",
	  "species_id": 411,
	  "height": 13,
	  "weight": 1495,
	  "base_experience": 173,
	  "order": 516,
	  "is_default": true
	},
	{
	  "id": 412,
	  "identifier": "burmy",
	  "species_id": 412,
	  "height": 2,
	  "weight": 34,
	  "base_experience": 45,
	  "order": 517,
	  "is_default": true
	},
	{
	  "id": 413,
	  "identifier": "wormadam-plant",
	  "species_id": 413,
	  "height": 5,
	  "weight": 65,
	  "base_experience": 148,
	  "order": 518,
	  "is_default": true
	},
	{
	  "id": 414,
	  "identifier": "mothim",
	  "species_id": 414,
	  "height": 9,
	  "weight": 233,
	  "base_experience": 148,
	  "order": 521,
	  "is_default": true
	},
	{
	  "id": 415,
	  "identifier": "combee",
	  "species_id": 415,
	  "height": 3,
	  "weight": 55,
	  "base_experience": 49,
	  "order": 522,
	  "is_default": true
	},
	{
	  "id": 416,
	  "identifier": "vespiquen",
	  "species_id": 416,
	  "height": 12,
	  "weight": 385,
	  "base_experience": 166,
	  "order": 523,
	  "is_default": true
	},
	{
	  "id": 417,
	  "identifier": "pachirisu",
	  "species_id": 417,
	  "height": 4,
	  "weight": 39,
	  "base_experience": 142,
	  "order": 524,
	  "is_default": true
	},
	{
	  "id": 418,
	  "identifier": "buizel",
	  "species_id": 418,
	  "height": 7,
	  "weight": 295,
	  "base_experience": 66,
	  "order": 525,
	  "is_default": true
	},
	{
	  "id": 419,
	  "identifier": "floatzel",
	  "species_id": 419,
	  "height": 11,
	  "weight": 335,
	  "base_experience": 173,
	  "order": 526,
	  "is_default": true
	},
	{
	  "id": 420,
	  "identifier": "cherubi",
	  "species_id": 420,
	  "height": 4,
	  "weight": 33,
	  "base_experience": 55,
	  "order": 527,
	  "is_default": true
	},
	{
	  "id": 421,
	  "identifier": "cherrim",
	  "species_id": 421,
	  "height": 5,
	  "weight": 93,
	  "base_experience": 158,
	  "order": 528,
	  "is_default": true
	},
	{
	  "id": 422,
	  "identifier": "shellos",
	  "species_id": 422,
	  "height": 3,
	  "weight": 63,
	  "base_experience": 65,
	  "order": 529,
	  "is_default": true
	},
	{
	  "id": 423,
	  "identifier": "gastrodon",
	  "species_id": 423,
	  "height": 9,
	  "weight": 299,
	  "base_experience": 166,
	  "order": 530,
	  "is_default": true
	},
	{
	  "id": 424,
	  "identifier": "ambipom",
	  "species_id": 424,
	  "height": 12,
	  "weight": 203,
	  "base_experience": 169,
	  "order": 265,
	  "is_default": true
	},
	{
	  "id": 425,
	  "identifier": "drifloon",
	  "species_id": 425,
	  "height": 4,
	  "weight": 12,
	  "base_experience": 70,
	  "order": 531,
	  "is_default": true
	},
	{
	  "id": 426,
	  "identifier": "drifblim",
	  "species_id": 426,
	  "height": 12,
	  "weight": 150,
	  "base_experience": 174,
	  "order": 532,
	  "is_default": true
	},
	{
	  "id": 427,
	  "identifier": "buneary",
	  "species_id": 427,
	  "height": 4,
	  "weight": 55,
	  "base_experience": 70,
	  "order": 533,
	  "is_default": true
	},
	{
	  "id": 428,
	  "identifier": "lopunny",
	  "species_id": 428,
	  "height": 12,
	  "weight": 333,
	  "base_experience": 168,
	  "order": 534,
	  "is_default": true
	},
	{
	  "id": 429,
	  "identifier": "mismagius",
	  "species_id": 429,
	  "height": 9,
	  "weight": 44,
	  "base_experience": 173,
	  "order": 275,
	  "is_default": true
	},
	{
	  "id": 430,
	  "identifier": "honchkrow",
	  "species_id": 430,
	  "height": 9,
	  "weight": 273,
	  "base_experience": 177,
	  "order": 273,
	  "is_default": true
	},
	{
	  "id": 431,
	  "identifier": "glameow",
	  "species_id": 431,
	  "height": 5,
	  "weight": 39,
	  "base_experience": 62,
	  "order": 536,
	  "is_default": true
	},
	{
	  "id": 432,
	  "identifier": "purugly",
	  "species_id": 432,
	  "height": 10,
	  "weight": 438,
	  "base_experience": 158,
	  "order": 537,
	  "is_default": true
	},
	{
	  "id": 433,
	  "identifier": "chingling",
	  "species_id": 433,
	  "height": 2,
	  "weight": 6,
	  "base_experience": 57,
	  "order": 452,
	  "is_default": true
	},
	{
	  "id": 434,
	  "identifier": "stunky",
	  "species_id": 434,
	  "height": 4,
	  "weight": 192,
	  "base_experience": 66,
	  "order": 538,
	  "is_default": true
	},
	{
	  "id": 435,
	  "identifier": "skuntank",
	  "species_id": 435,
	  "height": 10,
	  "weight": 380,
	  "base_experience": 168,
	  "order": 539,
	  "is_default": true
	},
	{
	  "id": 436,
	  "identifier": "bronzor",
	  "species_id": 436,
	  "height": 5,
	  "weight": 605,
	  "base_experience": 60,
	  "order": 540,
	  "is_default": true
	},
	{
	  "id": 437,
	  "identifier": "bronzong",
	  "species_id": 437,
	  "height": 13,
	  "weight": 1870,
	  "base_experience": 175,
	  "order": 541,
	  "is_default": true
	},
	{
	  "id": 438,
	  "identifier": "bonsly",
	  "species_id": 438,
	  "height": 5,
	  "weight": 150,
	  "base_experience": 58,
	  "order": 259,
	  "is_default": true
	},
	{
	  "id": 439,
	  "identifier": "mime-jr",
	  "species_id": 439,
	  "height": 6,
	  "weight": 130,
	  "base_experience": 62,
	  "order": 177,
	  "is_default": true
	},
	{
	  "id": 440,
	  "identifier": "happiny",
	  "species_id": 440,
	  "height": 6,
	  "weight": 244,
	  "base_experience": 110,
	  "order": 163,
	  "is_default": true
	},
	{
	  "id": 441,
	  "identifier": "chatot",
	  "species_id": 441,
	  "height": 5,
	  "weight": 19,
	  "base_experience": 144,
	  "order": 542,
	  "is_default": true
	},
	{
	  "id": 442,
	  "identifier": "spiritomb",
	  "species_id": 442,
	  "height": 10,
	  "weight": 1080,
	  "base_experience": 170,
	  "order": 543,
	  "is_default": true
	},
	{
	  "id": 443,
	  "identifier": "gible",
	  "species_id": 443,
	  "height": 7,
	  "weight": 205,
	  "base_experience": 60,
	  "order": 544,
	  "is_default": true
	},
	{
	  "id": 444,
	  "identifier": "gabite",
	  "species_id": 444,
	  "height": 14,
	  "weight": 560,
	  "base_experience": 144,
	  "order": 545,
	  "is_default": true
	},
	{
	  "id": 445,
	  "identifier": "garchomp",
	  "species_id": 445,
	  "height": 19,
	  "weight": 950,
	  "base_experience": 270,
	  "order": 546,
	  "is_default": true
	},
	{
	  "id": 446,
	  "identifier": "munchlax",
	  "species_id": 446,
	  "height": 6,
	  "weight": 1050,
	  "base_experience": 78,
	  "order": 216,
	  "is_default": true
	},
	{
	  "id": 447,
	  "identifier": "riolu",
	  "species_id": 447,
	  "height": 7,
	  "weight": 202,
	  "base_experience": 57,
	  "order": 548,
	  "is_default": true
	},
	{
	  "id": 448,
	  "identifier": "lucario",
	  "species_id": 448,
	  "height": 12,
	  "weight": 540,
	  "base_experience": 184,
	  "order": 549,
	  "is_default": true
	},
	{
	  "id": 449,
	  "identifier": "hippopotas",
	  "species_id": 449,
	  "height": 8,
	  "weight": 495,
	  "base_experience": 66,
	  "order": 551,
	  "is_default": true
	},
	{
	  "id": 450,
	  "identifier": "hippowdon",
	  "species_id": 450,
	  "height": 20,
	  "weight": 3000,
	  "base_experience": 184,
	  "order": 552,
	  "is_default": true
	},
	{
	  "id": 451,
	  "identifier": "skorupi",
	  "species_id": 451,
	  "height": 8,
	  "weight": 120,
	  "base_experience": 66,
	  "order": 553,
	  "is_default": true
	},
	{
	  "id": 452,
	  "identifier": "drapion",
	  "species_id": 452,
	  "height": 13,
	  "weight": 615,
	  "base_experience": 175,
	  "order": 554,
	  "is_default": true
	},
	{
	  "id": 453,
	  "identifier": "croagunk",
	  "species_id": 453,
	  "height": 7,
	  "weight": 230,
	  "base_experience": 60,
	  "order": 555,
	  "is_default": true
	},
	{
	  "id": 454,
	  "identifier": "toxicroak",
	  "species_id": 454,
	  "height": 13,
	  "weight": 444,
	  "base_experience": 172,
	  "order": 556,
	  "is_default": true
	},
	{
	  "id": 455,
	  "identifier": "carnivine",
	  "species_id": 455,
	  "height": 14,
	  "weight": 270,
	  "base_experience": 159,
	  "order": 557,
	  "is_default": true
	},
	{
	  "id": 456,
	  "identifier": "finneon",
	  "species_id": 456,
	  "height": 4,
	  "weight": 70,
	  "base_experience": 66,
	  "order": 558,
	  "is_default": true
	},
	{
	  "id": 457,
	  "identifier": "lumineon",
	  "species_id": 457,
	  "height": 12,
	  "weight": 240,
	  "base_experience": 161,
	  "order": 559,
	  "is_default": true
	},
	{
	  "id": 458,
	  "identifier": "mantyke",
	  "species_id": 458,
	  "height": 10,
	  "weight": 650,
	  "base_experience": 69,
	  "order": 304,
	  "is_default": true
	},
	{
	  "id": 459,
	  "identifier": "snover",
	  "species_id": 459,
	  "height": 10,
	  "weight": 505,
	  "base_experience": 67,
	  "order": 560,
	  "is_default": true
	},
	{
	  "id": 460,
	  "identifier": "abomasnow",
	  "species_id": 460,
	  "height": 22,
	  "weight": 1355,
	  "base_experience": 173,
	  "order": 561,
	  "is_default": true
	},
	{
	  "id": 461,
	  "identifier": "weavile",
	  "species_id": 461,
	  "height": 11,
	  "weight": 340,
	  "base_experience": 179,
	  "order": 292,
	  "is_default": true
	},
	{
	  "id": 462,
	  "identifier": "magnezone",
	  "species_id": 462,
	  "height": 12,
	  "weight": 1800,
	  "base_experience": 241,
	  "order": 120,
	  "is_default": true
	},
	{
	  "id": 463,
	  "identifier": "lickilicky",
	  "species_id": 463,
	  "height": 17,
	  "weight": 1400,
	  "base_experience": 180,
	  "order": 157,
	  "is_default": true
	},
	{
	  "id": 464,
	  "identifier": "rhyperior",
	  "species_id": 464,
	  "height": 24,
	  "weight": 2828,
	  "base_experience": 241,
	  "order": 162,
	  "is_default": true
	},
	{
	  "id": 465,
	  "identifier": "tangrowth",
	  "species_id": 465,
	  "height": 20,
	  "weight": 1286,
	  "base_experience": 187,
	  "order": 167,
	  "is_default": true
	},
	{
	  "id": 466,
	  "identifier": "electivire",
	  "species_id": 466,
	  "height": 18,
	  "weight": 1386,
	  "base_experience": 243,
	  "order": 186,
	  "is_default": true
	},
	{
	  "id": 467,
	  "identifier": "magmortar",
	  "species_id": 467,
	  "height": 16,
	  "weight": 680,
	  "base_experience": 243,
	  "order": 189,
	  "is_default": true
	},
	{
	  "id": 468,
	  "identifier": "togekiss",
	  "species_id": 468,
	  "height": 15,
	  "weight": 380,
	  "base_experience": 245,
	  "order": 249,
	  "is_default": true
	},
	{
	  "id": 469,
	  "identifier": "yanmega",
	  "species_id": 469,
	  "height": 19,
	  "weight": 515,
	  "base_experience": 180,
	  "order": 269,
	  "is_default": true
	},
	{
	  "id": 470,
	  "identifier": "leafeon",
	  "species_id": 470,
	  "height": 10,
	  "weight": 255,
	  "base_experience": 184,
	  "order": 204,
	  "is_default": true
	},
	{
	  "id": 471,
	  "identifier": "glaceon",
	  "species_id": 471,
	  "height": 8,
	  "weight": 259,
	  "base_experience": 184,
	  "order": 205,
	  "is_default": true
	},
	{
	  "id": 472,
	  "identifier": "gliscor",
	  "species_id": 472,
	  "height": 20,
	  "weight": 425,
	  "base_experience": 179,
	  "order": 284,
	  "is_default": true
	},
	{
	  "id": 473,
	  "identifier": "mamoswine",
	  "species_id": 473,
	  "height": 25,
	  "weight": 2910,
	  "base_experience": 239,
	  "order": 299,
	  "is_default": true
	},
	{
	  "id": 474,
	  "identifier": "porygon-z",
	  "species_id": 474,
	  "height": 9,
	  "weight": 340,
	  "base_experience": 241,
	  "order": 209,
	  "is_default": true
	},
	{
	  "id": 475,
	  "identifier": "gallade",
	  "species_id": 475,
	  "height": 16,
	  "weight": 520,
	  "base_experience": 233,
	  "order": 360,
	  "is_default": true
	},
	{
	  "id": 476,
	  "identifier": "probopass",
	  "species_id": 476,
	  "height": 14,
	  "weight": 3400,
	  "base_experience": 184,
	  "order": 378,
	  "is_default": true
	},
	{
	  "id": 477,
	  "identifier": "dusknoir",
	  "species_id": 477,
	  "height": 22,
	  "weight": 1066,
	  "base_experience": 236,
	  "order": 450,
	  "is_default": true
	},
	{
	  "id": 478,
	  "identifier": "froslass",
	  "species_id": 478,
	  "height": 13,
	  "weight": 266,
	  "base_experience": 168,
	  "order": 459,
	  "is_default": true
	},
	{
	  "id": 479,
	  "identifier": "rotom",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 154,
	  "order": 563,
	  "is_default": true
	},
	{
	  "id": 480,
	  "identifier": "uxie",
	  "species_id": 480,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 261,
	  "order": 569,
	  "is_default": true
	},
	{
	  "id": 481,
	  "identifier": "mesprit",
	  "species_id": 481,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 261,
	  "order": 570,
	  "is_default": true
	},
	{
	  "id": 482,
	  "identifier": "azelf",
	  "species_id": 482,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 261,
	  "order": 571,
	  "is_default": true
	},
	{
	  "id": 483,
	  "identifier": "dialga",
	  "species_id": 483,
	  "height": 54,
	  "weight": 6830,
	  "base_experience": 306,
	  "order": 572,
	  "is_default": true
	},
	{
	  "id": 484,
	  "identifier": "palkia",
	  "species_id": 484,
	  "height": 42,
	  "weight": 3360,
	  "base_experience": 306,
	  "order": 573,
	  "is_default": true
	},
	{
	  "id": 485,
	  "identifier": "heatran",
	  "species_id": 485,
	  "height": 17,
	  "weight": 4300,
	  "base_experience": 270,
	  "order": 574,
	  "is_default": true
	},
	{
	  "id": 486,
	  "identifier": "regigigas",
	  "species_id": 486,
	  "height": 37,
	  "weight": 4200,
	  "base_experience": 302,
	  "order": 575,
	  "is_default": true
	},
	{
	  "id": 487,
	  "identifier": "giratina-altered",
	  "species_id": 487,
	  "height": 45,
	  "weight": 7500,
	  "base_experience": 306,
	  "order": 576,
	  "is_default": true
	},
	{
	  "id": 488,
	  "identifier": "cresselia",
	  "species_id": 488,
	  "height": 15,
	  "weight": 856,
	  "base_experience": 270,
	  "order": 578,
	  "is_default": true
	},
	{
	  "id": 489,
	  "identifier": "phione",
	  "species_id": 489,
	  "height": 4,
	  "weight": 31,
	  "base_experience": 216,
	  "order": 579,
	  "is_default": true
	},
	{
	  "id": 490,
	  "identifier": "manaphy",
	  "species_id": 490,
	  "height": 3,
	  "weight": 14,
	  "base_experience": 270,
	  "order": 580,
	  "is_default": true
	},
	{
	  "id": 491,
	  "identifier": "darkrai",
	  "species_id": 491,
	  "height": 15,
	  "weight": 505,
	  "base_experience": 270,
	  "order": 581,
	  "is_default": true
	},
	{
	  "id": 492,
	  "identifier": "shaymin-land",
	  "species_id": 492,
	  "height": 2,
	  "weight": 21,
	  "base_experience": 270,
	  "order": 582,
	  "is_default": true
	},
	{
	  "id": 493,
	  "identifier": "arceus",
	  "species_id": 493,
	  "height": 32,
	  "weight": 3200,
	  "base_experience": 324,
	  "order": 584,
	  "is_default": true
	},
	{
	  "id": 494,
	  "identifier": "victini",
	  "species_id": 494,
	  "height": 4,
	  "weight": 40,
	  "base_experience": 270,
	  "order": 585,
	  "is_default": true
	},
	{
	  "id": 495,
	  "identifier": "snivy",
	  "species_id": 495,
	  "height": 6,
	  "weight": 81,
	  "base_experience": 62,
	  "order": 586,
	  "is_default": true
	},
	{
	  "id": 496,
	  "identifier": "servine",
	  "species_id": 496,
	  "height": 8,
	  "weight": 160,
	  "base_experience": 145,
	  "order": 587,
	  "is_default": true
	},
	{
	  "id": 497,
	  "identifier": "serperior",
	  "species_id": 497,
	  "height": 33,
	  "weight": 630,
	  "base_experience": 238,
	  "order": 588,
	  "is_default": true
	},
	{
	  "id": 498,
	  "identifier": "tepig",
	  "species_id": 498,
	  "height": 5,
	  "weight": 99,
	  "base_experience": 62,
	  "order": 589,
	  "is_default": true
	},
	{
	  "id": 499,
	  "identifier": "pignite",
	  "species_id": 499,
	  "height": 10,
	  "weight": 555,
	  "base_experience": 146,
	  "order": 590,
	  "is_default": true
	},
	{
	  "id": 500,
	  "identifier": "emboar",
	  "species_id": 500,
	  "height": 16,
	  "weight": 1500,
	  "base_experience": 238,
	  "order": 591,
	  "is_default": true
	},
	{
	  "id": 501,
	  "identifier": "oshawott",
	  "species_id": 501,
	  "height": 5,
	  "weight": 59,
	  "base_experience": 62,
	  "order": 592,
	  "is_default": true
	},
	{
	  "id": 502,
	  "identifier": "dewott",
	  "species_id": 502,
	  "height": 8,
	  "weight": 245,
	  "base_experience": 145,
	  "order": 593,
	  "is_default": true
	},
	{
	  "id": 503,
	  "identifier": "samurott",
	  "species_id": 503,
	  "height": 15,
	  "weight": 946,
	  "base_experience": 238,
	  "order": 594,
	  "is_default": true
	},
	{
	  "id": 504,
	  "identifier": "patrat",
	  "species_id": 504,
	  "height": 5,
	  "weight": 116,
	  "base_experience": 51,
	  "order": 595,
	  "is_default": true
	},
	{
	  "id": 505,
	  "identifier": "watchog",
	  "species_id": 505,
	  "height": 11,
	  "weight": 270,
	  "base_experience": 147,
	  "order": 596,
	  "is_default": true
	},
	{
	  "id": 506,
	  "identifier": "lillipup",
	  "species_id": 506,
	  "height": 4,
	  "weight": 41,
	  "base_experience": 55,
	  "order": 597,
	  "is_default": true
	},
	{
	  "id": 507,
	  "identifier": "herdier",
	  "species_id": 507,
	  "height": 9,
	  "weight": 147,
	  "base_experience": 130,
	  "order": 598,
	  "is_default": true
	},
	{
	  "id": 508,
	  "identifier": "stoutland",
	  "species_id": 508,
	  "height": 12,
	  "weight": 610,
	  "base_experience": 225,
	  "order": 599,
	  "is_default": true
	},
	{
	  "id": 509,
	  "identifier": "purrloin",
	  "species_id": 509,
	  "height": 4,
	  "weight": 101,
	  "base_experience": 56,
	  "order": 600,
	  "is_default": true
	},
	{
	  "id": 510,
	  "identifier": "liepard",
	  "species_id": 510,
	  "height": 11,
	  "weight": 375,
	  "base_experience": 156,
	  "order": 601,
	  "is_default": true
	},
	{
	  "id": 511,
	  "identifier": "pansage",
	  "species_id": 511,
	  "height": 6,
	  "weight": 105,
	  "base_experience": 63,
	  "order": 602,
	  "is_default": true
	},
	{
	  "id": 512,
	  "identifier": "simisage",
	  "species_id": 512,
	  "height": 11,
	  "weight": 305,
	  "base_experience": 174,
	  "order": 603,
	  "is_default": true
	},
	{
	  "id": 513,
	  "identifier": "pansear",
	  "species_id": 513,
	  "height": 6,
	  "weight": 110,
	  "base_experience": 63,
	  "order": 604,
	  "is_default": true
	},
	{
	  "id": 514,
	  "identifier": "simisear",
	  "species_id": 514,
	  "height": 10,
	  "weight": 280,
	  "base_experience": 174,
	  "order": 605,
	  "is_default": true
	},
	{
	  "id": 515,
	  "identifier": "panpour",
	  "species_id": 515,
	  "height": 6,
	  "weight": 135,
	  "base_experience": 63,
	  "order": 606,
	  "is_default": true
	},
	{
	  "id": 516,
	  "identifier": "simipour",
	  "species_id": 516,
	  "height": 10,
	  "weight": 290,
	  "base_experience": 174,
	  "order": 607,
	  "is_default": true
	},
	{
	  "id": 517,
	  "identifier": "munna",
	  "species_id": 517,
	  "height": 6,
	  "weight": 233,
	  "base_experience": 58,
	  "order": 608,
	  "is_default": true
	},
	{
	  "id": 518,
	  "identifier": "musharna",
	  "species_id": 518,
	  "height": 11,
	  "weight": 605,
	  "base_experience": 170,
	  "order": 609,
	  "is_default": true
	},
	{
	  "id": 519,
	  "identifier": "pidove",
	  "species_id": 519,
	  "height": 3,
	  "weight": 21,
	  "base_experience": 53,
	  "order": 610,
	  "is_default": true
	},
	{
	  "id": 520,
	  "identifier": "tranquill",
	  "species_id": 520,
	  "height": 6,
	  "weight": 150,
	  "base_experience": 125,
	  "order": 611,
	  "is_default": true
	},
	{
	  "id": 521,
	  "identifier": "unfezant",
	  "species_id": 521,
	  "height": 12,
	  "weight": 290,
	  "base_experience": 220,
	  "order": 612,
	  "is_default": true
	},
	{
	  "id": 522,
	  "identifier": "blitzle",
	  "species_id": 522,
	  "height": 8,
	  "weight": 298,
	  "base_experience": 59,
	  "order": 613,
	  "is_default": true
	},
	{
	  "id": 523,
	  "identifier": "zebstrika",
	  "species_id": 523,
	  "height": 16,
	  "weight": 795,
	  "base_experience": 174,
	  "order": 614,
	  "is_default": true
	},
	{
	  "id": 524,
	  "identifier": "roggenrola",
	  "species_id": 524,
	  "height": 4,
	  "weight": 180,
	  "base_experience": 56,
	  "order": 615,
	  "is_default": true
	},
	{
	  "id": 525,
	  "identifier": "boldore",
	  "species_id": 525,
	  "height": 9,
	  "weight": 1020,
	  "base_experience": 137,
	  "order": 616,
	  "is_default": true
	},
	{
	  "id": 526,
	  "identifier": "gigalith",
	  "species_id": 526,
	  "height": 17,
	  "weight": 2600,
	  "base_experience": 232,
	  "order": 617,
	  "is_default": true
	},
	{
	  "id": 527,
	  "identifier": "woobat",
	  "species_id": 527,
	  "height": 4,
	  "weight": 21,
	  "base_experience": 65,
	  "order": 618,
	  "is_default": true
	},
	{
	  "id": 528,
	  "identifier": "swoobat",
	  "species_id": 528,
	  "height": 9,
	  "weight": 105,
	  "base_experience": 149,
	  "order": 619,
	  "is_default": true
	},
	{
	  "id": 529,
	  "identifier": "drilbur",
	  "species_id": 529,
	  "height": 3,
	  "weight": 85,
	  "base_experience": 66,
	  "order": 620,
	  "is_default": true
	},
	{
	  "id": 530,
	  "identifier": "excadrill",
	  "species_id": 530,
	  "height": 7,
	  "weight": 404,
	  "base_experience": 178,
	  "order": 621,
	  "is_default": true
	},
	{
	  "id": 531,
	  "identifier": "audino",
	  "species_id": 531,
	  "height": 11,
	  "weight": 310,
	  "base_experience": 390,
	  "order": 622,
	  "is_default": true
	},
	{
	  "id": 532,
	  "identifier": "timburr",
	  "species_id": 532,
	  "height": 6,
	  "weight": 125,
	  "base_experience": 61,
	  "order": 624,
	  "is_default": true
	},
	{
	  "id": 533,
	  "identifier": "gurdurr",
	  "species_id": 533,
	  "height": 12,
	  "weight": 400,
	  "base_experience": 142,
	  "order": 625,
	  "is_default": true
	},
	{
	  "id": 534,
	  "identifier": "conkeldurr",
	  "species_id": 534,
	  "height": 14,
	  "weight": 870,
	  "base_experience": 227,
	  "order": 626,
	  "is_default": true
	},
	{
	  "id": 535,
	  "identifier": "tympole",
	  "species_id": 535,
	  "height": 5,
	  "weight": 45,
	  "base_experience": 59,
	  "order": 627,
	  "is_default": true
	},
	{
	  "id": 536,
	  "identifier": "palpitoad",
	  "species_id": 536,
	  "height": 8,
	  "weight": 170,
	  "base_experience": 134,
	  "order": 628,
	  "is_default": true
	},
	{
	  "id": 537,
	  "identifier": "seismitoad",
	  "species_id": 537,
	  "height": 15,
	  "weight": 620,
	  "base_experience": 229,
	  "order": 629,
	  "is_default": true
	},
	{
	  "id": 538,
	  "identifier": "throh",
	  "species_id": 538,
	  "height": 13,
	  "weight": 555,
	  "base_experience": 163,
	  "order": 630,
	  "is_default": true
	},
	{
	  "id": 539,
	  "identifier": "sawk",
	  "species_id": 539,
	  "height": 14,
	  "weight": 510,
	  "base_experience": 163,
	  "order": 631,
	  "is_default": true
	},
	{
	  "id": 540,
	  "identifier": "sewaddle",
	  "species_id": 540,
	  "height": 3,
	  "weight": 25,
	  "base_experience": 62,
	  "order": 632,
	  "is_default": true
	},
	{
	  "id": 541,
	  "identifier": "swadloon",
	  "species_id": 541,
	  "height": 5,
	  "weight": 73,
	  "base_experience": 133,
	  "order": 633,
	  "is_default": true
	},
	{
	  "id": 542,
	  "identifier": "leavanny",
	  "species_id": 542,
	  "height": 12,
	  "weight": 205,
	  "base_experience": 225,
	  "order": 634,
	  "is_default": true
	},
	{
	  "id": 543,
	  "identifier": "venipede",
	  "species_id": 543,
	  "height": 4,
	  "weight": 53,
	  "base_experience": 52,
	  "order": 635,
	  "is_default": true
	},
	{
	  "id": 544,
	  "identifier": "whirlipede",
	  "species_id": 544,
	  "height": 12,
	  "weight": 585,
	  "base_experience": 126,
	  "order": 636,
	  "is_default": true
	},
	{
	  "id": 545,
	  "identifier": "scolipede",
	  "species_id": 545,
	  "height": 25,
	  "weight": 2005,
	  "base_experience": 218,
	  "order": 637,
	  "is_default": true
	},
	{
	  "id": 546,
	  "identifier": "cottonee",
	  "species_id": 546,
	  "height": 3,
	  "weight": 6,
	  "base_experience": 56,
	  "order": 638,
	  "is_default": true
	},
	{
	  "id": 547,
	  "identifier": "whimsicott",
	  "species_id": 547,
	  "height": 7,
	  "weight": 66,
	  "base_experience": 168,
	  "order": 639,
	  "is_default": true
	},
	{
	  "id": 548,
	  "identifier": "petilil",
	  "species_id": 548,
	  "height": 5,
	  "weight": 66,
	  "base_experience": 56,
	  "order": 640,
	  "is_default": true
	},
	{
	  "id": 549,
	  "identifier": "lilligant",
	  "species_id": 549,
	  "height": 11,
	  "weight": 163,
	  "base_experience": 168,
	  "order": 641,
	  "is_default": true
	},
	{
	  "id": 550,
	  "identifier": "basculin-red-striped",
	  "species_id": 550,
	  "height": 10,
	  "weight": 180,
	  "base_experience": 161,
	  "order": 642,
	  "is_default": true
	},
	{
	  "id": 551,
	  "identifier": "sandile",
	  "species_id": 551,
	  "height": 7,
	  "weight": 152,
	  "base_experience": 58,
	  "order": 644,
	  "is_default": true
	},
	{
	  "id": 552,
	  "identifier": "krokorok",
	  "species_id": 552,
	  "height": 10,
	  "weight": 334,
	  "base_experience": 123,
	  "order": 645,
	  "is_default": true
	},
	{
	  "id": 553,
	  "identifier": "krookodile",
	  "species_id": 553,
	  "height": 15,
	  "weight": 963,
	  "base_experience": 234,
	  "order": 646,
	  "is_default": true
	},
	{
	  "id": 554,
	  "identifier": "darumaka",
	  "species_id": 554,
	  "height": 6,
	  "weight": 375,
	  "base_experience": 63,
	  "order": 647,
	  "is_default": true
	},
	{
	  "id": 555,
	  "identifier": "darmanitan-standard",
	  "species_id": 555,
	  "height": 13,
	  "weight": 929,
	  "base_experience": 168,
	  "order": 648,
	  "is_default": true
	},
	{
	  "id": 556,
	  "identifier": "maractus",
	  "species_id": 556,
	  "height": 10,
	  "weight": 280,
	  "base_experience": 161,
	  "order": 650,
	  "is_default": true
	},
	{
	  "id": 557,
	  "identifier": "dwebble",
	  "species_id": 557,
	  "height": 3,
	  "weight": 145,
	  "base_experience": 65,
	  "order": 651,
	  "is_default": true
	},
	{
	  "id": 558,
	  "identifier": "crustle",
	  "species_id": 558,
	  "height": 14,
	  "weight": 2000,
	  "base_experience": 170,
	  "order": 652,
	  "is_default": true
	},
	{
	  "id": 559,
	  "identifier": "scraggy",
	  "species_id": 559,
	  "height": 6,
	  "weight": 118,
	  "base_experience": 70,
	  "order": 653,
	  "is_default": true
	},
	{
	  "id": 560,
	  "identifier": "scrafty",
	  "species_id": 560,
	  "height": 11,
	  "weight": 300,
	  "base_experience": 171,
	  "order": 654,
	  "is_default": true
	},
	{
	  "id": 561,
	  "identifier": "sigilyph",
	  "species_id": 561,
	  "height": 14,
	  "weight": 140,
	  "base_experience": 172,
	  "order": 655,
	  "is_default": true
	},
	{
	  "id": 562,
	  "identifier": "yamask",
	  "species_id": 562,
	  "height": 5,
	  "weight": 15,
	  "base_experience": 61,
	  "order": 656,
	  "is_default": true
	},
	{
	  "id": 563,
	  "identifier": "cofagrigus",
	  "species_id": 563,
	  "height": 17,
	  "weight": 765,
	  "base_experience": 169,
	  "order": 657,
	  "is_default": true
	},
	{
	  "id": 564,
	  "identifier": "tirtouga",
	  "species_id": 564,
	  "height": 7,
	  "weight": 165,
	  "base_experience": 71,
	  "order": 658,
	  "is_default": true
	},
	{
	  "id": 565,
	  "identifier": "carracosta",
	  "species_id": 565,
	  "height": 12,
	  "weight": 810,
	  "base_experience": 173,
	  "order": 659,
	  "is_default": true
	},
	{
	  "id": 566,
	  "identifier": "archen",
	  "species_id": 566,
	  "height": 5,
	  "weight": 95,
	  "base_experience": 71,
	  "order": 660,
	  "is_default": true
	},
	{
	  "id": 567,
	  "identifier": "archeops",
	  "species_id": 567,
	  "height": 14,
	  "weight": 320,
	  "base_experience": 177,
	  "order": 661,
	  "is_default": true
	},
	{
	  "id": 568,
	  "identifier": "trubbish",
	  "species_id": 568,
	  "height": 6,
	  "weight": 310,
	  "base_experience": 66,
	  "order": 662,
	  "is_default": true
	},
	{
	  "id": 569,
	  "identifier": "garbodor",
	  "species_id": 569,
	  "height": 19,
	  "weight": 1073,
	  "base_experience": 166,
	  "order": 663,
	  "is_default": true
	},
	{
	  "id": 570,
	  "identifier": "zorua",
	  "species_id": 570,
	  "height": 7,
	  "weight": 125,
	  "base_experience": 66,
	  "order": 664,
	  "is_default": true
	},
	{
	  "id": 571,
	  "identifier": "zoroark",
	  "species_id": 571,
	  "height": 16,
	  "weight": 811,
	  "base_experience": 179,
	  "order": 665,
	  "is_default": true
	},
	{
	  "id": 572,
	  "identifier": "minccino",
	  "species_id": 572,
	  "height": 4,
	  "weight": 58,
	  "base_experience": 60,
	  "order": 666,
	  "is_default": true
	},
	{
	  "id": 573,
	  "identifier": "cinccino",
	  "species_id": 573,
	  "height": 5,
	  "weight": 75,
	  "base_experience": 165,
	  "order": 667,
	  "is_default": true
	},
	{
	  "id": 574,
	  "identifier": "gothita",
	  "species_id": 574,
	  "height": 4,
	  "weight": 58,
	  "base_experience": 58,
	  "order": 668,
	  "is_default": true
	},
	{
	  "id": 575,
	  "identifier": "gothorita",
	  "species_id": 575,
	  "height": 7,
	  "weight": 180,
	  "base_experience": 137,
	  "order": 669,
	  "is_default": true
	},
	{
	  "id": 576,
	  "identifier": "gothitelle",
	  "species_id": 576,
	  "height": 15,
	  "weight": 440,
	  "base_experience": 221,
	  "order": 670,
	  "is_default": true
	},
	{
	  "id": 577,
	  "identifier": "solosis",
	  "species_id": 577,
	  "height": 3,
	  "weight": 10,
	  "base_experience": 58,
	  "order": 671,
	  "is_default": true
	},
	{
	  "id": 578,
	  "identifier": "duosion",
	  "species_id": 578,
	  "height": 6,
	  "weight": 80,
	  "base_experience": 130,
	  "order": 672,
	  "is_default": true
	},
	{
	  "id": 579,
	  "identifier": "reuniclus",
	  "species_id": 579,
	  "height": 10,
	  "weight": 201,
	  "base_experience": 221,
	  "order": 673,
	  "is_default": true
	},
	{
	  "id": 580,
	  "identifier": "ducklett",
	  "species_id": 580,
	  "height": 5,
	  "weight": 55,
	  "base_experience": 61,
	  "order": 674,
	  "is_default": true
	},
	{
	  "id": 581,
	  "identifier": "swanna",
	  "species_id": 581,
	  "height": 13,
	  "weight": 242,
	  "base_experience": 166,
	  "order": 675,
	  "is_default": true
	},
	{
	  "id": 582,
	  "identifier": "vanillite",
	  "species_id": 582,
	  "height": 4,
	  "weight": 57,
	  "base_experience": 61,
	  "order": 676,
	  "is_default": true
	},
	{
	  "id": 583,
	  "identifier": "vanillish",
	  "species_id": 583,
	  "height": 11,
	  "weight": 410,
	  "base_experience": 138,
	  "order": 677,
	  "is_default": true
	},
	{
	  "id": 584,
	  "identifier": "vanilluxe",
	  "species_id": 584,
	  "height": 13,
	  "weight": 575,
	  "base_experience": 241,
	  "order": 678,
	  "is_default": true
	},
	{
	  "id": 585,
	  "identifier": "deerling",
	  "species_id": 585,
	  "height": 6,
	  "weight": 195,
	  "base_experience": 67,
	  "order": 679,
	  "is_default": true
	},
	{
	  "id": 586,
	  "identifier": "sawsbuck",
	  "species_id": 586,
	  "height": 19,
	  "weight": 925,
	  "base_experience": 166,
	  "order": 680,
	  "is_default": true
	},
	{
	  "id": 587,
	  "identifier": "emolga",
	  "species_id": 587,
	  "height": 4,
	  "weight": 50,
	  "base_experience": 150,
	  "order": 681,
	  "is_default": true
	},
	{
	  "id": 588,
	  "identifier": "karrablast",
	  "species_id": 588,
	  "height": 5,
	  "weight": 59,
	  "base_experience": 63,
	  "order": 682,
	  "is_default": true
	},
	{
	  "id": 589,
	  "identifier": "escavalier",
	  "species_id": 589,
	  "height": 10,
	  "weight": 330,
	  "base_experience": 173,
	  "order": 683,
	  "is_default": true
	},
	{
	  "id": 590,
	  "identifier": "foongus",
	  "species_id": 590,
	  "height": 2,
	  "weight": 10,
	  "base_experience": 59,
	  "order": 684,
	  "is_default": true
	},
	{
	  "id": 591,
	  "identifier": "amoonguss",
	  "species_id": 591,
	  "height": 6,
	  "weight": 105,
	  "base_experience": 162,
	  "order": 685,
	  "is_default": true
	},
	{
	  "id": 592,
	  "identifier": "frillish",
	  "species_id": 592,
	  "height": 12,
	  "weight": 330,
	  "base_experience": 67,
	  "order": 686,
	  "is_default": true
	},
	{
	  "id": 593,
	  "identifier": "jellicent",
	  "species_id": 593,
	  "height": 22,
	  "weight": 1350,
	  "base_experience": 168,
	  "order": 687,
	  "is_default": true
	},
	{
	  "id": 594,
	  "identifier": "alomomola",
	  "species_id": 594,
	  "height": 12,
	  "weight": 316,
	  "base_experience": 165,
	  "order": 688,
	  "is_default": true
	},
	{
	  "id": 595,
	  "identifier": "joltik",
	  "species_id": 595,
	  "height": 1,
	  "weight": 6,
	  "base_experience": 64,
	  "order": 689,
	  "is_default": true
	},
	{
	  "id": 596,
	  "identifier": "galvantula",
	  "species_id": 596,
	  "height": 8,
	  "weight": 143,
	  "base_experience": 165,
	  "order": 690,
	  "is_default": true
	},
	{
	  "id": 597,
	  "identifier": "ferroseed",
	  "species_id": 597,
	  "height": 6,
	  "weight": 188,
	  "base_experience": 61,
	  "order": 691,
	  "is_default": true
	},
	{
	  "id": 598,
	  "identifier": "ferrothorn",
	  "species_id": 598,
	  "height": 10,
	  "weight": 1100,
	  "base_experience": 171,
	  "order": 692,
	  "is_default": true
	},
	{
	  "id": 599,
	  "identifier": "klink",
	  "species_id": 599,
	  "height": 3,
	  "weight": 210,
	  "base_experience": 60,
	  "order": 693,
	  "is_default": true
	},
	{
	  "id": 600,
	  "identifier": "klang",
	  "species_id": 600,
	  "height": 6,
	  "weight": 510,
	  "base_experience": 154,
	  "order": 694,
	  "is_default": true
	},
	{
	  "id": 601,
	  "identifier": "klinklang",
	  "species_id": 601,
	  "height": 6,
	  "weight": 810,
	  "base_experience": 234,
	  "order": 695,
	  "is_default": true
	},
	{
	  "id": 602,
	  "identifier": "tynamo",
	  "species_id": 602,
	  "height": 2,
	  "weight": 3,
	  "base_experience": 55,
	  "order": 696,
	  "is_default": true
	},
	{
	  "id": 603,
	  "identifier": "eelektrik",
	  "species_id": 603,
	  "height": 12,
	  "weight": 220,
	  "base_experience": 142,
	  "order": 697,
	  "is_default": true
	},
	{
	  "id": 604,
	  "identifier": "eelektross",
	  "species_id": 604,
	  "height": 21,
	  "weight": 805,
	  "base_experience": 232,
	  "order": 698,
	  "is_default": true
	},
	{
	  "id": 605,
	  "identifier": "elgyem",
	  "species_id": 605,
	  "height": 5,
	  "weight": 90,
	  "base_experience": 67,
	  "order": 699,
	  "is_default": true
	},
	{
	  "id": 606,
	  "identifier": "beheeyem",
	  "species_id": 606,
	  "height": 10,
	  "weight": 345,
	  "base_experience": 170,
	  "order": 700,
	  "is_default": true
	},
	{
	  "id": 607,
	  "identifier": "litwick",
	  "species_id": 607,
	  "height": 3,
	  "weight": 31,
	  "base_experience": 55,
	  "order": 701,
	  "is_default": true
	},
	{
	  "id": 608,
	  "identifier": "lampent",
	  "species_id": 608,
	  "height": 6,
	  "weight": 130,
	  "base_experience": 130,
	  "order": 702,
	  "is_default": true
	},
	{
	  "id": 609,
	  "identifier": "chandelure",
	  "species_id": 609,
	  "height": 10,
	  "weight": 343,
	  "base_experience": 234,
	  "order": 703,
	  "is_default": true
	},
	{
	  "id": 610,
	  "identifier": "axew",
	  "species_id": 610,
	  "height": 6,
	  "weight": 180,
	  "base_experience": 64,
	  "order": 704,
	  "is_default": true
	},
	{
	  "id": 611,
	  "identifier": "fraxure",
	  "species_id": 611,
	  "height": 10,
	  "weight": 360,
	  "base_experience": 144,
	  "order": 705,
	  "is_default": true
	},
	{
	  "id": 612,
	  "identifier": "haxorus",
	  "species_id": 612,
	  "height": 18,
	  "weight": 1055,
	  "base_experience": 243,
	  "order": 706,
	  "is_default": true
	},
	{
	  "id": 613,
	  "identifier": "cubchoo",
	  "species_id": 613,
	  "height": 5,
	  "weight": 85,
	  "base_experience": 61,
	  "order": 707,
	  "is_default": true
	},
	{
	  "id": 614,
	  "identifier": "beartic",
	  "species_id": 614,
	  "height": 26,
	  "weight": 2600,
	  "base_experience": 177,
	  "order": 708,
	  "is_default": true
	},
	{
	  "id": 615,
	  "identifier": "cryogonal",
	  "species_id": 615,
	  "height": 11,
	  "weight": 1480,
	  "base_experience": 180,
	  "order": 709,
	  "is_default": true
	},
	{
	  "id": 616,
	  "identifier": "shelmet",
	  "species_id": 616,
	  "height": 4,
	  "weight": 77,
	  "base_experience": 61,
	  "order": 710,
	  "is_default": true
	},
	{
	  "id": 617,
	  "identifier": "accelgor",
	  "species_id": 617,
	  "height": 8,
	  "weight": 253,
	  "base_experience": 173,
	  "order": 711,
	  "is_default": true
	},
	{
	  "id": 618,
	  "identifier": "stunfisk",
	  "species_id": 618,
	  "height": 7,
	  "weight": 110,
	  "base_experience": 165,
	  "order": 712,
	  "is_default": true
	},
	{
	  "id": 619,
	  "identifier": "mienfoo",
	  "species_id": 619,
	  "height": 9,
	  "weight": 200,
	  "base_experience": 70,
	  "order": 713,
	  "is_default": true
	},
	{
	  "id": 620,
	  "identifier": "mienshao",
	  "species_id": 620,
	  "height": 14,
	  "weight": 355,
	  "base_experience": 179,
	  "order": 714,
	  "is_default": true
	},
	{
	  "id": 621,
	  "identifier": "druddigon",
	  "species_id": 621,
	  "height": 16,
	  "weight": 1390,
	  "base_experience": 170,
	  "order": 715,
	  "is_default": true
	},
	{
	  "id": 622,
	  "identifier": "golett",
	  "species_id": 622,
	  "height": 10,
	  "weight": 920,
	  "base_experience": 61,
	  "order": 716,
	  "is_default": true
	},
	{
	  "id": 623,
	  "identifier": "golurk",
	  "species_id": 623,
	  "height": 28,
	  "weight": 3300,
	  "base_experience": 169,
	  "order": 717,
	  "is_default": true
	},
	{
	  "id": 624,
	  "identifier": "pawniard",
	  "species_id": 624,
	  "height": 5,
	  "weight": 102,
	  "base_experience": 68,
	  "order": 718,
	  "is_default": true
	},
	{
	  "id": 625,
	  "identifier": "bisharp",
	  "species_id": 625,
	  "height": 16,
	  "weight": 700,
	  "base_experience": 172,
	  "order": 719,
	  "is_default": true
	},
	{
	  "id": 626,
	  "identifier": "bouffalant",
	  "species_id": 626,
	  "height": 16,
	  "weight": 946,
	  "base_experience": 172,
	  "order": 720,
	  "is_default": true
	},
	{
	  "id": 627,
	  "identifier": "rufflet",
	  "species_id": 627,
	  "height": 5,
	  "weight": 105,
	  "base_experience": 70,
	  "order": 721,
	  "is_default": true
	},
	{
	  "id": 628,
	  "identifier": "braviary",
	  "species_id": 628,
	  "height": 15,
	  "weight": 410,
	  "base_experience": 179,
	  "order": 722,
	  "is_default": true
	},
	{
	  "id": 629,
	  "identifier": "vullaby",
	  "species_id": 629,
	  "height": 5,
	  "weight": 90,
	  "base_experience": 74,
	  "order": 723,
	  "is_default": true
	},
	{
	  "id": 630,
	  "identifier": "mandibuzz",
	  "species_id": 630,
	  "height": 12,
	  "weight": 395,
	  "base_experience": 179,
	  "order": 724,
	  "is_default": true
	},
	{
	  "id": 631,
	  "identifier": "heatmor",
	  "species_id": 631,
	  "height": 14,
	  "weight": 580,
	  "base_experience": 169,
	  "order": 725,
	  "is_default": true
	},
	{
	  "id": 632,
	  "identifier": "durant",
	  "species_id": 632,
	  "height": 3,
	  "weight": 330,
	  "base_experience": 169,
	  "order": 726,
	  "is_default": true
	},
	{
	  "id": 633,
	  "identifier": "deino",
	  "species_id": 633,
	  "height": 8,
	  "weight": 173,
	  "base_experience": 60,
	  "order": 727,
	  "is_default": true
	},
	{
	  "id": 634,
	  "identifier": "zweilous",
	  "species_id": 634,
	  "height": 14,
	  "weight": 500,
	  "base_experience": 147,
	  "order": 728,
	  "is_default": true
	},
	{
	  "id": 635,
	  "identifier": "hydreigon",
	  "species_id": 635,
	  "height": 18,
	  "weight": 1600,
	  "base_experience": 270,
	  "order": 729,
	  "is_default": true
	},
	{
	  "id": 636,
	  "identifier": "larvesta",
	  "species_id": 636,
	  "height": 11,
	  "weight": 288,
	  "base_experience": 72,
	  "order": 730,
	  "is_default": true
	},
	{
	  "id": 637,
	  "identifier": "volcarona",
	  "species_id": 637,
	  "height": 16,
	  "weight": 460,
	  "base_experience": 248,
	  "order": 731,
	  "is_default": true
	},
	{
	  "id": 638,
	  "identifier": "cobalion",
	  "species_id": 638,
	  "height": 21,
	  "weight": 2500,
	  "base_experience": 261,
	  "order": 732,
	  "is_default": true
	},
	{
	  "id": 639,
	  "identifier": "terrakion",
	  "species_id": 639,
	  "height": 19,
	  "weight": 2600,
	  "base_experience": 261,
	  "order": 733,
	  "is_default": true
	},
	{
	  "id": 640,
	  "identifier": "virizion",
	  "species_id": 640,
	  "height": 20,
	  "weight": 2000,
	  "base_experience": 261,
	  "order": 734,
	  "is_default": true
	},
	{
	  "id": 641,
	  "identifier": "tornadus-incarnate",
	  "species_id": 641,
	  "height": 15,
	  "weight": 630,
	  "base_experience": 261,
	  "order": 735,
	  "is_default": true
	},
	{
	  "id": 642,
	  "identifier": "thundurus-incarnate",
	  "species_id": 642,
	  "height": 15,
	  "weight": 610,
	  "base_experience": 261,
	  "order": 737,
	  "is_default": true
	},
	{
	  "id": 643,
	  "identifier": "reshiram",
	  "species_id": 643,
	  "height": 32,
	  "weight": 3300,
	  "base_experience": 306,
	  "order": 739,
	  "is_default": true
	},
	{
	  "id": 644,
	  "identifier": "zekrom",
	  "species_id": 644,
	  "height": 29,
	  "weight": 3450,
	  "base_experience": 306,
	  "order": 740,
	  "is_default": true
	},
	{
	  "id": 645,
	  "identifier": "landorus-incarnate",
	  "species_id": 645,
	  "height": 15,
	  "weight": 680,
	  "base_experience": 270,
	  "order": 741,
	  "is_default": true
	},
	{
	  "id": 646,
	  "identifier": "kyurem",
	  "species_id": 646,
	  "height": 30,
	  "weight": 3250,
	  "base_experience": 297,
	  "order": 743,
	  "is_default": true
	},
	{
	  "id": 647,
	  "identifier": "keldeo-ordinary",
	  "species_id": 647,
	  "height": 14,
	  "weight": 485,
	  "base_experience": 261,
	  "order": 746,
	  "is_default": true
	},
	{
	  "id": 648,
	  "identifier": "meloetta-aria",
	  "species_id": 648,
	  "height": 6,
	  "weight": 65,
	  "base_experience": 270,
	  "order": 748,
	  "is_default": true
	},
	{
	  "id": 649,
	  "identifier": "genesect",
	  "species_id": 649,
	  "height": 15,
	  "weight": 825,
	  "base_experience": 270,
	  "order": 750,
	  "is_default": true
	},
	{
	  "id": 650,
	  "identifier": "chespin",
	  "species_id": 650,
	  "height": 4,
	  "weight": 90,
	  "base_experience": 63,
	  "order": 751,
	  "is_default": true
	},
	{
	  "id": 651,
	  "identifier": "quilladin",
	  "species_id": 651,
	  "height": 7,
	  "weight": 290,
	  "base_experience": 142,
	  "order": 752,
	  "is_default": true
	},
	{
	  "id": 652,
	  "identifier": "chesnaught",
	  "species_id": 652,
	  "height": 16,
	  "weight": 900,
	  "base_experience": 239,
	  "order": 753,
	  "is_default": true
	},
	{
	  "id": 653,
	  "identifier": "fennekin",
	  "species_id": 653,
	  "height": 4,
	  "weight": 94,
	  "base_experience": 61,
	  "order": 754,
	  "is_default": true
	},
	{
	  "id": 654,
	  "identifier": "braixen",
	  "species_id": 654,
	  "height": 10,
	  "weight": 145,
	  "base_experience": 143,
	  "order": 755,
	  "is_default": true
	},
	{
	  "id": 655,
	  "identifier": "delphox",
	  "species_id": 655,
	  "height": 15,
	  "weight": 390,
	  "base_experience": 240,
	  "order": 756,
	  "is_default": true
	},
	{
	  "id": 656,
	  "identifier": "froakie",
	  "species_id": 656,
	  "height": 3,
	  "weight": 70,
	  "base_experience": 63,
	  "order": 757,
	  "is_default": true
	},
	{
	  "id": 657,
	  "identifier": "frogadier",
	  "species_id": 657,
	  "height": 6,
	  "weight": 109,
	  "base_experience": 142,
	  "order": 758,
	  "is_default": true
	},
	{
	  "id": 658,
	  "identifier": "greninja",
	  "species_id": 658,
	  "height": 15,
	  "weight": 400,
	  "base_experience": 239,
	  "order": 759,
	  "is_default": true
	},
	{
	  "id": 659,
	  "identifier": "bunnelby",
	  "species_id": 659,
	  "height": 4,
	  "weight": 50,
	  "base_experience": 47,
	  "order": 762,
	  "is_default": true
	},
	{
	  "id": 660,
	  "identifier": "diggersby",
	  "species_id": 660,
	  "height": 10,
	  "weight": 424,
	  "base_experience": 148,
	  "order": 763,
	  "is_default": true
	},
	{
	  "id": 661,
	  "identifier": "fletchling",
	  "species_id": 661,
	  "height": 3,
	  "weight": 17,
	  "base_experience": 56,
	  "order": 764,
	  "is_default": true
	},
	{
	  "id": 662,
	  "identifier": "fletchinder",
	  "species_id": 662,
	  "height": 7,
	  "weight": 160,
	  "base_experience": 134,
	  "order": 765,
	  "is_default": true
	},
	{
	  "id": 663,
	  "identifier": "talonflame",
	  "species_id": 663,
	  "height": 12,
	  "weight": 245,
	  "base_experience": 175,
	  "order": 766,
	  "is_default": true
	},
	{
	  "id": 664,
	  "identifier": "scatterbug",
	  "species_id": 664,
	  "height": 3,
	  "weight": 25,
	  "base_experience": 40,
	  "order": 767,
	  "is_default": true
	},
	{
	  "id": 665,
	  "identifier": "spewpa",
	  "species_id": 665,
	  "height": 3,
	  "weight": 84,
	  "base_experience": 75,
	  "order": 768,
	  "is_default": true
	},
	{
	  "id": 666,
	  "identifier": "vivillon",
	  "species_id": 666,
	  "height": 12,
	  "weight": 170,
	  "base_experience": 185,
	  "order": 769,
	  "is_default": true
	},
	{
	  "id": 667,
	  "identifier": "litleo",
	  "species_id": 667,
	  "height": 6,
	  "weight": 135,
	  "base_experience": 74,
	  "order": 770,
	  "is_default": true
	},
	{
	  "id": 668,
	  "identifier": "pyroar",
	  "species_id": 668,
	  "height": 15,
	  "weight": 815,
	  "base_experience": 177,
	  "order": 771,
	  "is_default": true
	},
	{
	  "id": 669,
	  "identifier": "flabebe",
	  "species_id": 669,
	  "height": 1,
	  "weight": 1,
	  "base_experience": 61,
	  "order": 772,
	  "is_default": true
	},
	{
	  "id": 670,
	  "identifier": "floette",
	  "species_id": 670,
	  "height": 2,
	  "weight": 9,
	  "base_experience": 130,
	  "order": 773,
	  "is_default": true
	},
	{
	  "id": 671,
	  "identifier": "florges",
	  "species_id": 671,
	  "height": 11,
	  "weight": 100,
	  "base_experience": 248,
	  "order": 775,
	  "is_default": true
	},
	{
	  "id": 672,
	  "identifier": "skiddo",
	  "species_id": 672,
	  "height": 9,
	  "weight": 310,
	  "base_experience": 70,
	  "order": 776,
	  "is_default": true
	},
	{
	  "id": 673,
	  "identifier": "gogoat",
	  "species_id": 673,
	  "height": 17,
	  "weight": 910,
	  "base_experience": 186,
	  "order": 777,
	  "is_default": true
	},
	{
	  "id": 674,
	  "identifier": "pancham",
	  "species_id": 674,
	  "height": 6,
	  "weight": 80,
	  "base_experience": 70,
	  "order": 778,
	  "is_default": true
	},
	{
	  "id": 675,
	  "identifier": "pangoro",
	  "species_id": 675,
	  "height": 21,
	  "weight": 1360,
	  "base_experience": 173,
	  "order": 779,
	  "is_default": true
	},
	{
	  "id": 676,
	  "identifier": "furfrou",
	  "species_id": 676,
	  "height": 12,
	  "weight": 280,
	  "base_experience": 165,
	  "order": 780,
	  "is_default": true
	},
	{
	  "id": 677,
	  "identifier": "espurr",
	  "species_id": 677,
	  "height": 3,
	  "weight": 35,
	  "base_experience": 71,
	  "order": 781,
	  "is_default": true
	},
	{
	  "id": 678,
	  "identifier": "meowstic-male",
	  "species_id": 678,
	  "height": 6,
	  "weight": 85,
	  "base_experience": 163,
	  "order": 782,
	  "is_default": true
	},
	{
	  "id": 679,
	  "identifier": "honedge",
	  "species_id": 679,
	  "height": 8,
	  "weight": 20,
	  "base_experience": 65,
	  "order": 784,
	  "is_default": true
	},
	{
	  "id": 680,
	  "identifier": "doublade",
	  "species_id": 680,
	  "height": 8,
	  "weight": 45,
	  "base_experience": 157,
	  "order": 785,
	  "is_default": true
	},
	{
	  "id": 681,
	  "identifier": "aegislash-shield",
	  "species_id": 681,
	  "height": 17,
	  "weight": 530,
	  "base_experience": 234,
	  "order": 786,
	  "is_default": true
	},
	{
	  "id": 682,
	  "identifier": "spritzee",
	  "species_id": 682,
	  "height": 2,
	  "weight": 5,
	  "base_experience": 68,
	  "order": 788,
	  "is_default": true
	},
	{
	  "id": 683,
	  "identifier": "aromatisse",
	  "species_id": 683,
	  "height": 8,
	  "weight": 155,
	  "base_experience": 162,
	  "order": 789,
	  "is_default": true
	},
	{
	  "id": 684,
	  "identifier": "swirlix",
	  "species_id": 684,
	  "height": 4,
	  "weight": 35,
	  "base_experience": 68,
	  "order": 790,
	  "is_default": true
	},
	{
	  "id": 685,
	  "identifier": "slurpuff",
	  "species_id": 685,
	  "height": 8,
	  "weight": 50,
	  "base_experience": 168,
	  "order": 791,
	  "is_default": true
	},
	{
	  "id": 686,
	  "identifier": "inkay",
	  "species_id": 686,
	  "height": 4,
	  "weight": 35,
	  "base_experience": 58,
	  "order": 792,
	  "is_default": true
	},
	{
	  "id": 687,
	  "identifier": "malamar",
	  "species_id": 687,
	  "height": 15,
	  "weight": 470,
	  "base_experience": 169,
	  "order": 793,
	  "is_default": true
	},
	{
	  "id": 688,
	  "identifier": "binacle",
	  "species_id": 688,
	  "height": 5,
	  "weight": 310,
	  "base_experience": 61,
	  "order": 794,
	  "is_default": true
	},
	{
	  "id": 689,
	  "identifier": "barbaracle",
	  "species_id": 689,
	  "height": 13,
	  "weight": 960,
	  "base_experience": 175,
	  "order": 795,
	  "is_default": true
	},
	{
	  "id": 690,
	  "identifier": "skrelp",
	  "species_id": 690,
	  "height": 5,
	  "weight": 73,
	  "base_experience": 64,
	  "order": 796,
	  "is_default": true
	},
	{
	  "id": 691,
	  "identifier": "dragalge",
	  "species_id": 691,
	  "height": 18,
	  "weight": 815,
	  "base_experience": 173,
	  "order": 797,
	  "is_default": true
	},
	{
	  "id": 692,
	  "identifier": "clauncher",
	  "species_id": 692,
	  "height": 5,
	  "weight": 83,
	  "base_experience": 66,
	  "order": 798,
	  "is_default": true
	},
	{
	  "id": 693,
	  "identifier": "clawitzer",
	  "species_id": 693,
	  "height": 13,
	  "weight": 353,
	  "base_experience": 100,
	  "order": 799,
	  "is_default": true
	},
	{
	  "id": 694,
	  "identifier": "helioptile",
	  "species_id": 694,
	  "height": 5,
	  "weight": 60,
	  "base_experience": 58,
	  "order": 800,
	  "is_default": true
	},
	{
	  "id": 695,
	  "identifier": "heliolisk",
	  "species_id": 695,
	  "height": 10,
	  "weight": 210,
	  "base_experience": 168,
	  "order": 801,
	  "is_default": true
	},
	{
	  "id": 696,
	  "identifier": "tyrunt",
	  "species_id": 696,
	  "height": 8,
	  "weight": 260,
	  "base_experience": 72,
	  "order": 802,
	  "is_default": true
	},
	{
	  "id": 697,
	  "identifier": "tyrantrum",
	  "species_id": 697,
	  "height": 25,
	  "weight": 2700,
	  "base_experience": 182,
	  "order": 803,
	  "is_default": true
	},
	{
	  "id": 698,
	  "identifier": "amaura",
	  "species_id": 698,
	  "height": 13,
	  "weight": 252,
	  "base_experience": 72,
	  "order": 804,
	  "is_default": true
	},
	{
	  "id": 699,
	  "identifier": "aurorus",
	  "species_id": 699,
	  "height": 27,
	  "weight": 2250,
	  "base_experience": 104,
	  "order": 805,
	  "is_default": true
	},
	{
	  "id": 700,
	  "identifier": "sylveon",
	  "species_id": 700,
	  "height": 10,
	  "weight": 235,
	  "base_experience": 184,
	  "order": 206,
	  "is_default": true
	},
	{
	  "id": 701,
	  "identifier": "hawlucha",
	  "species_id": 701,
	  "height": 8,
	  "weight": 215,
	  "base_experience": 175,
	  "order": 806,
	  "is_default": true
	},
	{
	  "id": 702,
	  "identifier": "dedenne",
	  "species_id": 702,
	  "height": 2,
	  "weight": 22,
	  "base_experience": 151,
	  "order": 807,
	  "is_default": true
	},
	{
	  "id": 703,
	  "identifier": "carbink",
	  "species_id": 703,
	  "height": 3,
	  "weight": 57,
	  "base_experience": 100,
	  "order": 808,
	  "is_default": true
	},
	{
	  "id": 704,
	  "identifier": "goomy",
	  "species_id": 704,
	  "height": 3,
	  "weight": 28,
	  "base_experience": 60,
	  "order": 809,
	  "is_default": true
	},
	{
	  "id": 705,
	  "identifier": "sliggoo",
	  "species_id": 705,
	  "height": 8,
	  "weight": 175,
	  "base_experience": 158,
	  "order": 810,
	  "is_default": true
	},
	{
	  "id": 706,
	  "identifier": "goodra",
	  "species_id": 706,
	  "height": 20,
	  "weight": 1505,
	  "base_experience": 270,
	  "order": 811,
	  "is_default": true
	},
	{
	  "id": 707,
	  "identifier": "klefki",
	  "species_id": 707,
	  "height": 2,
	  "weight": 30,
	  "base_experience": 165,
	  "order": 812,
	  "is_default": true
	},
	{
	  "id": 708,
	  "identifier": "phantump",
	  "species_id": 708,
	  "height": 4,
	  "weight": 70,
	  "base_experience": 62,
	  "order": 813,
	  "is_default": true
	},
	{
	  "id": 709,
	  "identifier": "trevenant",
	  "species_id": 709,
	  "height": 15,
	  "weight": 710,
	  "base_experience": 166,
	  "order": 814,
	  "is_default": true
	},
	{
	  "id": 710,
	  "identifier": "pumpkaboo-average",
	  "species_id": 710,
	  "height": 4,
	  "weight": 50,
	  "base_experience": 67,
	  "order": 815,
	  "is_default": true
	},
	{
	  "id": 711,
	  "identifier": "gourgeist-average",
	  "species_id": 711,
	  "height": 9,
	  "weight": 125,
	  "base_experience": 173,
	  "order": 819,
	  "is_default": true
	},
	{
	  "id": 712,
	  "identifier": "bergmite",
	  "species_id": 712,
	  "height": 10,
	  "weight": 995,
	  "base_experience": 61,
	  "order": 823,
	  "is_default": true
	},
	{
	  "id": 713,
	  "identifier": "avalugg",
	  "species_id": 713,
	  "height": 20,
	  "weight": 5050,
	  "base_experience": 180,
	  "order": 824,
	  "is_default": true
	},
	{
	  "id": 714,
	  "identifier": "noibat",
	  "species_id": 714,
	  "height": 5,
	  "weight": 80,
	  "base_experience": 49,
	  "order": 825,
	  "is_default": true
	},
	{
	  "id": 715,
	  "identifier": "noivern",
	  "species_id": 715,
	  "height": 15,
	  "weight": 850,
	  "base_experience": 187,
	  "order": 826,
	  "is_default": true
	},
	{
	  "id": 716,
	  "identifier": "xerneas",
	  "species_id": 716,
	  "height": 30,
	  "weight": 2150,
	  "base_experience": 306,
	  "order": 827,
	  "is_default": true
	},
	{
	  "id": 717,
	  "identifier": "yveltal",
	  "species_id": 717,
	  "height": 58,
	  "weight": 2030,
	  "base_experience": 306,
	  "order": 828,
	  "is_default": true
	},
	{
	  "id": 718,
	  "identifier": "zygarde",
	  "species_id": 718,
	  "height": 50,
	  "weight": 3050,
	  "base_experience": 270,
	  "order": 829,
	  "is_default": true
	},
	{
	  "id": 719,
	  "identifier": "diancie",
	  "species_id": 719,
	  "height": 7,
	  "weight": 88,
	  "base_experience": 270,
	  "order": 833,
	  "is_default": true
	},
	{
	  "id": 720,
	  "identifier": "hoopa",
	  "species_id": 720,
	  "height": 5,
	  "weight": 90,
	  "base_experience": 270,
	  "order": 835,
	  "is_default": true
	},
	{
	  "id": 721,
	  "identifier": "volcanion",
	  "species_id": 721,
	  "height": 17,
	  "weight": 1950,
	  "base_experience": 270,
	  "order": 837,
	  "is_default": true
	},
	{
	  "id": 722,
	  "identifier": "rowlet",
	  "species_id": 722,
	  "height": 3,
	  "weight": 15,
	  "base_experience": 64,
	  "order": 838,
	  "is_default": true
	},
	{
	  "id": 723,
	  "identifier": "dartrix",
	  "species_id": 723,
	  "height": 7,
	  "weight": 160,
	  "base_experience": 147,
	  "order": 839,
	  "is_default": true
	},
	{
	  "id": 724,
	  "identifier": "decidueye",
	  "species_id": 724,
	  "height": 16,
	  "weight": 366,
	  "base_experience": 239,
	  "order": 840,
	  "is_default": true
	},
	{
	  "id": 725,
	  "identifier": "litten",
	  "species_id": 725,
	  "height": 4,
	  "weight": 43,
	  "base_experience": 64,
	  "order": 841,
	  "is_default": true
	},
	{
	  "id": 726,
	  "identifier": "torracat",
	  "species_id": 726,
	  "height": 7,
	  "weight": 250,
	  "base_experience": 147,
	  "order": 842,
	  "is_default": true
	},
	{
	  "id": 727,
	  "identifier": "incineroar",
	  "species_id": 727,
	  "height": 18,
	  "weight": 830,
	  "base_experience": 239,
	  "order": 843,
	  "is_default": true
	},
	{
	  "id": 728,
	  "identifier": "popplio",
	  "species_id": 728,
	  "height": 4,
	  "weight": 75,
	  "base_experience": 64,
	  "order": 844,
	  "is_default": true
	},
	{
	  "id": 729,
	  "identifier": "brionne",
	  "species_id": 729,
	  "height": 6,
	  "weight": 175,
	  "base_experience": 147,
	  "order": 845,
	  "is_default": true
	},
	{
	  "id": 730,
	  "identifier": "primarina",
	  "species_id": 730,
	  "height": 18,
	  "weight": 440,
	  "base_experience": 239,
	  "order": 846,
	  "is_default": true
	},
	{
	  "id": 731,
	  "identifier": "pikipek",
	  "species_id": 731,
	  "height": 3,
	  "weight": 12,
	  "base_experience": 53,
	  "order": 847,
	  "is_default": true
	},
	{
	  "id": 732,
	  "identifier": "trumbeak",
	  "species_id": 732,
	  "height": 6,
	  "weight": 148,
	  "base_experience": 124,
	  "order": 848,
	  "is_default": true
	},
	{
	  "id": 733,
	  "identifier": "toucannon",
	  "species_id": 733,
	  "height": 11,
	  "weight": 260,
	  "base_experience": 218,
	  "order": 849,
	  "is_default": true
	},
	{
	  "id": 734,
	  "identifier": "yungoos",
	  "species_id": 734,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 51,
	  "order": 850,
	  "is_default": true
	},
	{
	  "id": 735,
	  "identifier": "gumshoos",
	  "species_id": 735,
	  "height": 7,
	  "weight": 142,
	  "base_experience": 146,
	  "order": 851,
	  "is_default": true
	},
	{
	  "id": 736,
	  "identifier": "grubbin",
	  "species_id": 736,
	  "height": 4,
	  "weight": 44,
	  "base_experience": 60,
	  "order": 853,
	  "is_default": true
	},
	{
	  "id": 737,
	  "identifier": "charjabug",
	  "species_id": 737,
	  "height": 5,
	  "weight": 105,
	  "base_experience": 140,
	  "order": 854,
	  "is_default": true
	},
	{
	  "id": 738,
	  "identifier": "vikavolt",
	  "species_id": 738,
	  "height": 15,
	  "weight": 450,
	  "base_experience": 225,
	  "order": 855,
	  "is_default": true
	},
	{
	  "id": 739,
	  "identifier": "crabrawler",
	  "species_id": 739,
	  "height": 6,
	  "weight": 70,
	  "base_experience": 68,
	  "order": 857,
	  "is_default": true
	},
	{
	  "id": 740,
	  "identifier": "crabominable",
	  "species_id": 740,
	  "height": 17,
	  "weight": 1800,
	  "base_experience": 167,
	  "order": 858,
	  "is_default": true
	},
	{
	  "id": 741,
	  "identifier": "oricorio-baile",
	  "species_id": 741,
	  "height": 6,
	  "weight": 34,
	  "base_experience": 167,
	  "order": 859,
	  "is_default": true
	},
	{
	  "id": 742,
	  "identifier": "cutiefly",
	  "species_id": 742,
	  "height": 1,
	  "weight": 2,
	  "base_experience": 61,
	  "order": 863,
	  "is_default": true
	},
	{
	  "id": 743,
	  "identifier": "ribombee",
	  "species_id": 743,
	  "height": 2,
	  "weight": 5,
	  "base_experience": 162,
	  "order": 864,
	  "is_default": true
	},
	{
	  "id": 744,
	  "identifier": "rockruff",
	  "species_id": 744,
	  "height": 5,
	  "weight": 92,
	  "base_experience": 56,
	  "order": 866,
	  "is_default": true
	},
	{
	  "id": 745,
	  "identifier": "lycanroc-midday",
	  "species_id": 745,
	  "height": 8,
	  "weight": 250,
	  "base_experience": 170,
	  "order": 868,
	  "is_default": true
	},
	{
	  "id": 746,
	  "identifier": "wishiwashi-solo",
	  "species_id": 746,
	  "height": 2,
	  "weight": 3,
	  "base_experience": 61,
	  "order": 871,
	  "is_default": true
	},
	{
	  "id": 747,
	  "identifier": "mareanie",
	  "species_id": 747,
	  "height": 4,
	  "weight": 80,
	  "base_experience": 61,
	  "order": 873,
	  "is_default": true
	},
	{
	  "id": 748,
	  "identifier": "toxapex",
	  "species_id": 748,
	  "height": 7,
	  "weight": 145,
	  "base_experience": 173,
	  "order": 874,
	  "is_default": true
	},
	{
	  "id": 749,
	  "identifier": "mudbray",
	  "species_id": 749,
	  "height": 10,
	  "weight": 1100,
	  "base_experience": 77,
	  "order": 875,
	  "is_default": true
	},
	{
	  "id": 750,
	  "identifier": "mudsdale",
	  "species_id": 750,
	  "height": 25,
	  "weight": 9200,
	  "base_experience": 175,
	  "order": 876,
	  "is_default": true
	},
	{
	  "id": 751,
	  "identifier": "dewpider",
	  "species_id": 751,
	  "height": 3,
	  "weight": 40,
	  "base_experience": 54,
	  "order": 877,
	  "is_default": true
	},
	{
	  "id": 752,
	  "identifier": "araquanid",
	  "species_id": 752,
	  "height": 18,
	  "weight": 820,
	  "base_experience": 159,
	  "order": 878,
	  "is_default": true
	},
	{
	  "id": 753,
	  "identifier": "fomantis",
	  "species_id": 753,
	  "height": 3,
	  "weight": 15,
	  "base_experience": 50,
	  "order": 880,
	  "is_default": true
	},
	{
	  "id": 754,
	  "identifier": "lurantis",
	  "species_id": 754,
	  "height": 9,
	  "weight": 185,
	  "base_experience": 168,
	  "order": 881,
	  "is_default": true
	},
	{
	  "id": 755,
	  "identifier": "morelull",
	  "species_id": 755,
	  "height": 2,
	  "weight": 15,
	  "base_experience": 57,
	  "order": 883,
	  "is_default": true
	},
	{
	  "id": 756,
	  "identifier": "shiinotic",
	  "species_id": 756,
	  "height": 10,
	  "weight": 115,
	  "base_experience": 142,
	  "order": 884,
	  "is_default": true
	},
	{
	  "id": 757,
	  "identifier": "salandit",
	  "species_id": 757,
	  "height": 6,
	  "weight": 48,
	  "base_experience": 64,
	  "order": 885,
	  "is_default": true
	},
	{
	  "id": 758,
	  "identifier": "salazzle",
	  "species_id": 758,
	  "height": 12,
	  "weight": 222,
	  "base_experience": 168,
	  "order": 886,
	  "is_default": true
	},
	{
	  "id": 759,
	  "identifier": "stufful",
	  "species_id": 759,
	  "height": 5,
	  "weight": 68,
	  "base_experience": 68,
	  "order": 888,
	  "is_default": true
	},
	{
	  "id": 760,
	  "identifier": "bewear",
	  "species_id": 760,
	  "height": 21,
	  "weight": 1350,
	  "base_experience": 175,
	  "order": 889,
	  "is_default": true
	},
	{
	  "id": 761,
	  "identifier": "bounsweet",
	  "species_id": 761,
	  "height": 3,
	  "weight": 32,
	  "base_experience": 42,
	  "order": 890,
	  "is_default": true
	},
	{
	  "id": 762,
	  "identifier": "steenee",
	  "species_id": 762,
	  "height": 7,
	  "weight": 82,
	  "base_experience": 102,
	  "order": 891,
	  "is_default": true
	},
	{
	  "id": 763,
	  "identifier": "tsareena",
	  "species_id": 763,
	  "height": 12,
	  "weight": 214,
	  "base_experience": 230,
	  "order": 892,
	  "is_default": true
	},
	{
	  "id": 764,
	  "identifier": "comfey",
	  "species_id": 764,
	  "height": 1,
	  "weight": 3,
	  "base_experience": 170,
	  "order": 893,
	  "is_default": true
	},
	{
	  "id": 765,
	  "identifier": "oranguru",
	  "species_id": 765,
	  "height": 15,
	  "weight": 760,
	  "base_experience": 172,
	  "order": 894,
	  "is_default": true
	},
	{
	  "id": 766,
	  "identifier": "passimian",
	  "species_id": 766,
	  "height": 20,
	  "weight": 828,
	  "base_experience": 172,
	  "order": 895,
	  "is_default": true
	},
	{
	  "id": 767,
	  "identifier": "wimpod",
	  "species_id": 767,
	  "height": 5,
	  "weight": 120,
	  "base_experience": 46,
	  "order": 896,
	  "is_default": true
	},
	{
	  "id": 768,
	  "identifier": "golisopod",
	  "species_id": 768,
	  "height": 20,
	  "weight": 1080,
	  "base_experience": 186,
	  "order": 897,
	  "is_default": true
	},
	{
	  "id": 769,
	  "identifier": "sandygast",
	  "species_id": 769,
	  "height": 5,
	  "weight": 700,
	  "base_experience": 64,
	  "order": 898,
	  "is_default": true
	},
	{
	  "id": 770,
	  "identifier": "palossand",
	  "species_id": 770,
	  "height": 13,
	  "weight": 2500,
	  "base_experience": 168,
	  "order": 899,
	  "is_default": true
	},
	{
	  "id": 771,
	  "identifier": "pyukumuku",
	  "species_id": 771,
	  "height": 3,
	  "weight": 12,
	  "base_experience": 144,
	  "order": 900,
	  "is_default": true
	},
	{
	  "id": 772,
	  "identifier": "type-null",
	  "species_id": 772,
	  "height": 19,
	  "weight": 1205,
	  "base_experience": 107,
	  "order": 901,
	  "is_default": true
	},
	{
	  "id": 773,
	  "identifier": "silvally",
	  "species_id": 773,
	  "height": 23,
	  "weight": 1005,
	  "base_experience": 257,
	  "order": 902,
	  "is_default": true
	},
	{
	  "id": 774,
	  "identifier": "minior-red-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 903,
	  "is_default": true
	},
	{
	  "id": 775,
	  "identifier": "komala",
	  "species_id": 775,
	  "height": 4,
	  "weight": 199,
	  "base_experience": 168,
	  "order": 917,
	  "is_default": true
	},
	{
	  "id": 776,
	  "identifier": "turtonator",
	  "species_id": 776,
	  "height": 20,
	  "weight": 2120,
	  "base_experience": 170,
	  "order": 918,
	  "is_default": true
	},
	{
	  "id": 777,
	  "identifier": "togedemaru",
	  "species_id": 777,
	  "height": 3,
	  "weight": 33,
	  "base_experience": 152,
	  "order": 919,
	  "is_default": true
	},
	{
	  "id": 778,
	  "identifier": "mimikyu-disguised",
	  "species_id": 778,
	  "height": 2,
	  "weight": 7,
	  "base_experience": 167,
	  "order": 921,
	  "is_default": true
	},
	{
	  "id": 779,
	  "identifier": "bruxish",
	  "species_id": 779,
	  "height": 9,
	  "weight": 190,
	  "base_experience": 166,
	  "order": 925,
	  "is_default": true
	},
	{
	  "id": 780,
	  "identifier": "drampa",
	  "species_id": 780,
	  "height": 30,
	  "weight": 1850,
	  "base_experience": 170,
	  "order": 926,
	  "is_default": true
	},
	{
	  "id": 781,
	  "identifier": "dhelmise",
	  "species_id": 781,
	  "height": 39,
	  "weight": 2100,
	  "base_experience": 181,
	  "order": 927,
	  "is_default": true
	},
	{
	  "id": 782,
	  "identifier": "jangmo-o",
	  "species_id": 782,
	  "height": 6,
	  "weight": 297,
	  "base_experience": 60,
	  "order": 928,
	  "is_default": true
	},
	{
	  "id": 783,
	  "identifier": "hakamo-o",
	  "species_id": 783,
	  "height": 12,
	  "weight": 470,
	  "base_experience": 147,
	  "order": 929,
	  "is_default": true
	},
	{
	  "id": 784,
	  "identifier": "kommo-o",
	  "species_id": 784,
	  "height": 16,
	  "weight": 782,
	  "base_experience": 270,
	  "order": 930,
	  "is_default": true
	},
	{
	  "id": 785,
	  "identifier": "tapu-koko",
	  "species_id": 785,
	  "height": 18,
	  "weight": 205,
	  "base_experience": 257,
	  "order": 932,
	  "is_default": true
	},
	{
	  "id": 786,
	  "identifier": "tapu-lele",
	  "species_id": 786,
	  "height": 12,
	  "weight": 186,
	  "base_experience": 257,
	  "order": 933,
	  "is_default": true
	},
	{
	  "id": 787,
	  "identifier": "tapu-bulu",
	  "species_id": 787,
	  "height": 19,
	  "weight": 455,
	  "base_experience": 257,
	  "order": 934,
	  "is_default": true
	},
	{
	  "id": 788,
	  "identifier": "tapu-fini",
	  "species_id": 788,
	  "height": 13,
	  "weight": 212,
	  "base_experience": 257,
	  "order": 935,
	  "is_default": true
	},
	{
	  "id": 789,
	  "identifier": "cosmog",
	  "species_id": 789,
	  "height": 2,
	  "weight": 1,
	  "base_experience": 40,
	  "order": 936,
	  "is_default": true
	},
	{
	  "id": 790,
	  "identifier": "cosmoem",
	  "species_id": 790,
	  "height": 1,
	  "weight": 9999,
	  "base_experience": 140,
	  "order": 937,
	  "is_default": true
	},
	{
	  "id": 791,
	  "identifier": "solgaleo",
	  "species_id": 791,
	  "height": 34,
	  "weight": 2300,
	  "base_experience": 306,
	  "order": 938,
	  "is_default": true
	},
	{
	  "id": 792,
	  "identifier": "lunala",
	  "species_id": 792,
	  "height": 40,
	  "weight": 1200,
	  "base_experience": 306,
	  "order": 939,
	  "is_default": true
	},
	{
	  "id": 793,
	  "identifier": "nihilego",
	  "species_id": 793,
	  "height": 12,
	  "weight": 555,
	  "base_experience": 257,
	  "order": 940,
	  "is_default": true
	},
	{
	  "id": 794,
	  "identifier": "buzzwole",
	  "species_id": 794,
	  "height": 24,
	  "weight": 3336,
	  "base_experience": 257,
	  "order": 941,
	  "is_default": true
	},
	{
	  "id": 795,
	  "identifier": "pheromosa",
	  "species_id": 795,
	  "height": 18,
	  "weight": 250,
	  "base_experience": 257,
	  "order": 942,
	  "is_default": true
	},
	{
	  "id": 796,
	  "identifier": "xurkitree",
	  "species_id": 796,
	  "height": 38,
	  "weight": 1000,
	  "base_experience": 257,
	  "order": 943,
	  "is_default": true
	},
	{
	  "id": 797,
	  "identifier": "celesteela",
	  "species_id": 797,
	  "height": 92,
	  "weight": 9999,
	  "base_experience": 257,
	  "order": 944,
	  "is_default": true
	},
	{
	  "id": 798,
	  "identifier": "kartana",
	  "species_id": 798,
	  "height": 3,
	  "weight": 1,
	  "base_experience": 257,
	  "order": 945,
	  "is_default": true
	},
	{
	  "id": 799,
	  "identifier": "guzzlord",
	  "species_id": 799,
	  "height": 55,
	  "weight": 8880,
	  "base_experience": 257,
	  "order": 946,
	  "is_default": true
	},
	{
	  "id": 800,
	  "identifier": "necrozma",
	  "species_id": 800,
	  "height": 24,
	  "weight": 2300,
	  "base_experience": 270,
	  "order": 947,
	  "is_default": true
	},
	{
	  "id": 801,
	  "identifier": "magearna",
	  "species_id": 801,
	  "height": 10,
	  "weight": 805,
	  "base_experience": 270,
	  "order": 951,
	  "is_default": true
	},
	{
	  "id": 802,
	  "identifier": "marshadow",
	  "species_id": 802,
	  "height": 7,
	  "weight": 222,
	  "base_experience": 270,
	  "order": 953,
	  "is_default": true
	},
	{
	  "id": 803,
	  "identifier": "poipole",
	  "species_id": 803,
	  "height": 6,
	  "weight": 18,
	  "base_experience": 189,
	  "order": 954,
	  "is_default": true
	},
	{
	  "id": 804,
	  "identifier": "naganadel",
	  "species_id": 804,
	  "height": 36,
	  "weight": 1500,
	  "base_experience": 243,
	  "order": 955,
	  "is_default": true
	},
	{
	  "id": 805,
	  "identifier": "stakataka",
	  "species_id": 805,
	  "height": 55,
	  "weight": 8200,
	  "base_experience": 257,
	  "order": 956,
	  "is_default": true
	},
	{
	  "id": 806,
	  "identifier": "blacephalon",
	  "species_id": 806,
	  "height": 18,
	  "weight": 130,
	  "base_experience": 257,
	  "order": 957,
	  "is_default": true
	},
	{
	  "id": 807,
	  "identifier": "zeraora",
	  "species_id": 807,
	  "height": 15,
	  "weight": 445,
	  "base_experience": 270,
	  "order": 958,
	  "is_default": true
	},
	{
	  "id": 10001,
	  "identifier": "deoxys-attack",
	  "species_id": 386,
	  "height": 17,
	  "weight": 608,
	  "base_experience": 270,
	  "order": 491,
	  "is_default": false
	},
	{
	  "id": 10002,
	  "identifier": "deoxys-defense",
	  "species_id": 386,
	  "height": 17,
	  "weight": 608,
	  "base_experience": 270,
	  "order": 492,
	  "is_default": false
	},
	{
	  "id": 10003,
	  "identifier": "deoxys-speed",
	  "species_id": 386,
	  "height": 17,
	  "weight": 608,
	  "base_experience": 270,
	  "order": 493,
	  "is_default": false
	},
	{
	  "id": 10004,
	  "identifier": "wormadam-sandy",
	  "species_id": 413,
	  "height": 5,
	  "weight": 65,
	  "base_experience": 148,
	  "order": 519,
	  "is_default": false
	},
	{
	  "id": 10005,
	  "identifier": "wormadam-trash",
	  "species_id": 413,
	  "height": 5,
	  "weight": 65,
	  "base_experience": 148,
	  "order": 520,
	  "is_default": false
	},
	{
	  "id": 10006,
	  "identifier": "shaymin-sky",
	  "species_id": 492,
	  "height": 4,
	  "weight": 52,
	  "base_experience": 270,
	  "order": 583,
	  "is_default": false
	},
	{
	  "id": 10007,
	  "identifier": "giratina-origin",
	  "species_id": 487,
	  "height": 69,
	  "weight": 6500,
	  "base_experience": 306,
	  "order": 577,
	  "is_default": false
	},
	{
	  "id": 10008,
	  "identifier": "rotom-heat",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 182,
	  "order": 564,
	  "is_default": false
	},
	{
	  "id": 10009,
	  "identifier": "rotom-wash",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 182,
	  "order": 565,
	  "is_default": false
	},
	{
	  "id": 10010,
	  "identifier": "rotom-frost",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 182,
	  "order": 566,
	  "is_default": false
	},
	{
	  "id": 10011,
	  "identifier": "rotom-fan",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 182,
	  "order": 567,
	  "is_default": false
	},
	{
	  "id": 10012,
	  "identifier": "rotom-mow",
	  "species_id": 479,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 182,
	  "order": 568,
	  "is_default": false
	},
	{
	  "id": 10013,
	  "identifier": "castform-sunny",
	  "species_id": 351,
	  "height": 3,
	  "weight": 8,
	  "base_experience": 147,
	  "order": 441,
	  "is_default": false
	},
	{
	  "id": 10014,
	  "identifier": "castform-rainy",
	  "species_id": 351,
	  "height": 3,
	  "weight": 8,
	  "base_experience": 147,
	  "order": 442,
	  "is_default": false
	},
	{
	  "id": 10015,
	  "identifier": "castform-snowy",
	  "species_id": 351,
	  "height": 3,
	  "weight": 8,
	  "base_experience": 147,
	  "order": 443,
	  "is_default": false
	},
	{
	  "id": 10016,
	  "identifier": "basculin-blue-striped",
	  "species_id": 550,
	  "height": 10,
	  "weight": 180,
	  "base_experience": 161,
	  "order": 643,
	  "is_default": false
	},
	{
	  "id": 10017,
	  "identifier": "darmanitan-zen",
	  "species_id": 555,
	  "height": 13,
	  "weight": 929,
	  "base_experience": 189,
	  "order": 649,
	  "is_default": false
	},
	{
	  "id": 10018,
	  "identifier": "meloetta-pirouette",
	  "species_id": 648,
	  "height": 6,
	  "weight": 65,
	  "base_experience": 270,
	  "order": 749,
	  "is_default": false
	},
	{
	  "id": 10019,
	  "identifier": "tornadus-therian",
	  "species_id": 641,
	  "height": 14,
	  "weight": 630,
	  "base_experience": 261,
	  "order": 736,
	  "is_default": false
	},
	{
	  "id": 10020,
	  "identifier": "thundurus-therian",
	  "species_id": 642,
	  "height": 30,
	  "weight": 610,
	  "base_experience": 261,
	  "order": 738,
	  "is_default": false
	},
	{
	  "id": 10021,
	  "identifier": "landorus-therian",
	  "species_id": 645,
	  "height": 13,
	  "weight": 680,
	  "base_experience": 270,
	  "order": 742,
	  "is_default": false
	},
	{
	  "id": 10022,
	  "identifier": "kyurem-black",
	  "species_id": 646,
	  "height": 33,
	  "weight": 3250,
	  "base_experience": 315,
	  "order": 745,
	  "is_default": false
	},
	{
	  "id": 10023,
	  "identifier": "kyurem-white",
	  "species_id": 646,
	  "height": 36,
	  "weight": 3250,
	  "base_experience": 315,
	  "order": 744,
	  "is_default": false
	},
	{
	  "id": 10024,
	  "identifier": "keldeo-resolute",
	  "species_id": 647,
	  "height": 14,
	  "weight": 485,
	  "base_experience": 261,
	  "order": 747,
	  "is_default": false
	},
	{
	  "id": 10025,
	  "identifier": "meowstic-female",
	  "species_id": 678,
	  "height": 6,
	  "weight": 85,
	  "base_experience": 163,
	  "order": 783,
	  "is_default": false
	},
	{
	  "id": 10026,
	  "identifier": "aegislash-blade",
	  "species_id": 681,
	  "height": 17,
	  "weight": 530,
	  "base_experience": 234,
	  "order": 787,
	  "is_default": false
	},
	{
	  "id": 10027,
	  "identifier": "pumpkaboo-small",
	  "species_id": 710,
	  "height": 3,
	  "weight": 35,
	  "base_experience": 67,
	  "order": 816,
	  "is_default": false
	},
	{
	  "id": 10028,
	  "identifier": "pumpkaboo-large",
	  "species_id": 710,
	  "height": 5,
	  "weight": 75,
	  "base_experience": 67,
	  "order": 817,
	  "is_default": false
	},
	{
	  "id": 10029,
	  "identifier": "pumpkaboo-super",
	  "species_id": 710,
	  "height": 8,
	  "weight": 150,
	  "base_experience": 67,
	  "order": 818,
	  "is_default": false
	},
	{
	  "id": 10030,
	  "identifier": "gourgeist-small",
	  "species_id": 711,
	  "height": 7,
	  "weight": 95,
	  "base_experience": 173,
	  "order": 820,
	  "is_default": false
	},
	{
	  "id": 10031,
	  "identifier": "gourgeist-large",
	  "species_id": 711,
	  "height": 11,
	  "weight": 140,
	  "base_experience": 173,
	  "order": 821,
	  "is_default": false
	},
	{
	  "id": 10032,
	  "identifier": "gourgeist-super",
	  "species_id": 711,
	  "height": 17,
	  "weight": 390,
	  "base_experience": 173,
	  "order": 822,
	  "is_default": false
	},
	{
	  "id": 10033,
	  "identifier": "venusaur-mega",
	  "species_id": 3,
	  "height": 24,
	  "weight": 1555,
	  "base_experience": 281,
	  "order": 4,
	  "is_default": false
	},
	{
	  "id": 10034,
	  "identifier": "charizard-mega-x",
	  "species_id": 6,
	  "height": 17,
	  "weight": 1105,
	  "base_experience": 285,
	  "order": 8,
	  "is_default": false
	},
	{
	  "id": 10035,
	  "identifier": "charizard-mega-y",
	  "species_id": 6,
	  "height": 17,
	  "weight": 1005,
	  "base_experience": 285,
	  "order": 9,
	  "is_default": false
	},
	{
	  "id": 10036,
	  "identifier": "blastoise-mega",
	  "species_id": 9,
	  "height": 16,
	  "weight": 1011,
	  "base_experience": 284,
	  "order": 13,
	  "is_default": false
	},
	{
	  "id": 10037,
	  "identifier": "alakazam-mega",
	  "species_id": 65,
	  "height": 12,
	  "weight": 480,
	  "base_experience": 270,
	  "order": 97,
	  "is_default": false
	},
	{
	  "id": 10038,
	  "identifier": "gengar-mega",
	  "species_id": 94,
	  "height": 14,
	  "weight": 405,
	  "base_experience": 270,
	  "order": 135,
	  "is_default": false
	},
	{
	  "id": 10039,
	  "identifier": "kangaskhan-mega",
	  "species_id": 115,
	  "height": 22,
	  "weight": 1000,
	  "base_experience": 207,
	  "order": 169,
	  "is_default": false
	},
	{
	  "id": 10040,
	  "identifier": "pinsir-mega",
	  "species_id": 127,
	  "height": 17,
	  "weight": 590,
	  "base_experience": 210,
	  "order": 191,
	  "is_default": false
	},
	{
	  "id": 10041,
	  "identifier": "gyarados-mega",
	  "species_id": 130,
	  "height": 65,
	  "weight": 3050,
	  "base_experience": 224,
	  "order": 195,
	  "is_default": false
	},
	{
	  "id": 10042,
	  "identifier": "aerodactyl-mega",
	  "species_id": 142,
	  "height": 21,
	  "weight": 790,
	  "base_experience": 215,
	  "order": 215,
	  "is_default": false
	},
	{
	  "id": 10043,
	  "identifier": "mewtwo-mega-x",
	  "species_id": 150,
	  "height": 23,
	  "weight": 1270,
	  "base_experience": 351,
	  "order": 225,
	  "is_default": false
	},
	{
	  "id": 10044,
	  "identifier": "mewtwo-mega-y",
	  "species_id": 150,
	  "height": 15,
	  "weight": 330,
	  "base_experience": 351,
	  "order": 226,
	  "is_default": false
	},
	{
	  "id": 10045,
	  "identifier": "ampharos-mega",
	  "species_id": 181,
	  "height": 14,
	  "weight": 615,
	  "base_experience": 275,
	  "order": 255,
	  "is_default": false
	},
	{
	  "id": 10046,
	  "identifier": "scizor-mega",
	  "species_id": 212,
	  "height": 20,
	  "weight": 1250,
	  "base_experience": 210,
	  "order": 181,
	  "is_default": false
	},
	{
	  "id": 10047,
	  "identifier": "heracross-mega",
	  "species_id": 214,
	  "height": 17,
	  "weight": 625,
	  "base_experience": 210,
	  "order": 290,
	  "is_default": false
	},
	{
	  "id": 10048,
	  "identifier": "houndoom-mega",
	  "species_id": 229,
	  "height": 19,
	  "weight": 495,
	  "base_experience": 210,
	  "order": 309,
	  "is_default": false
	},
	{
	  "id": 10049,
	  "identifier": "tyranitar-mega",
	  "species_id": 248,
	  "height": 25,
	  "weight": 2550,
	  "base_experience": 315,
	  "order": 321,
	  "is_default": false
	},
	{
	  "id": 10050,
	  "identifier": "blaziken-mega",
	  "species_id": 257,
	  "height": 19,
	  "weight": 520,
	  "base_experience": 284,
	  "order": 332,
	  "is_default": false
	},
	{
	  "id": 10051,
	  "identifier": "gardevoir-mega",
	  "species_id": 282,
	  "height": 16,
	  "weight": 484,
	  "base_experience": 278,
	  "order": 359,
	  "is_default": false
	},
	{
	  "id": 10052,
	  "identifier": "mawile-mega",
	  "species_id": 303,
	  "height": 10,
	  "weight": 235,
	  "base_experience": 168,
	  "order": 384,
	  "is_default": false
	},
	{
	  "id": 10053,
	  "identifier": "aggron-mega",
	  "species_id": 306,
	  "height": 22,
	  "weight": 3950,
	  "base_experience": 284,
	  "order": 388,
	  "is_default": false
	},
	{
	  "id": 10054,
	  "identifier": "medicham-mega",
	  "species_id": 308,
	  "height": 13,
	  "weight": 315,
	  "base_experience": 179,
	  "order": 391,
	  "is_default": false
	},
	{
	  "id": 10055,
	  "identifier": "manectric-mega",
	  "species_id": 310,
	  "height": 18,
	  "weight": 440,
	  "base_experience": 201,
	  "order": 394,
	  "is_default": false
	},
	{
	  "id": 10056,
	  "identifier": "banette-mega",
	  "species_id": 354,
	  "height": 12,
	  "weight": 130,
	  "base_experience": 194,
	  "order": 447,
	  "is_default": false
	},
	{
	  "id": 10057,
	  "identifier": "absol-mega",
	  "species_id": 359,
	  "height": 12,
	  "weight": 490,
	  "base_experience": 198,
	  "order": 455,
	  "is_default": false
	},
	{
	  "id": 10058,
	  "identifier": "garchomp-mega",
	  "species_id": 445,
	  "height": 19,
	  "weight": 950,
	  "base_experience": 315,
	  "order": 547,
	  "is_default": false
	},
	{
	  "id": 10059,
	  "identifier": "lucario-mega",
	  "species_id": 448,
	  "height": 13,
	  "weight": 575,
	  "base_experience": 219,
	  "order": 550,
	  "is_default": false
	},
	{
	  "id": 10060,
	  "identifier": "abomasnow-mega",
	  "species_id": 460,
	  "height": 27,
	  "weight": 1850,
	  "base_experience": 208,
	  "order": 562,
	  "is_default": false
	},
	{
	  "id": 10061,
	  "identifier": "floette-eternal",
	  "species_id": 670,
	  "height": 2,
	  "weight": 9,
	  "base_experience": 243,
	  "order": 774,
	  "is_default": false
	},
	{
	  "id": 10062,
	  "identifier": "latias-mega",
	  "species_id": 380,
	  "height": 18,
	  "weight": 520,
	  "base_experience": 315,
	  "order": 480,
	  "is_default": false
	},
	{
	  "id": 10063,
	  "identifier": "latios-mega",
	  "species_id": 381,
	  "height": 23,
	  "weight": 700,
	  "base_experience": 315,
	  "order": 482,
	  "is_default": false
	},
	{
	  "id": 10064,
	  "identifier": "swampert-mega",
	  "species_id": 260,
	  "height": 19,
	  "weight": 1020,
	  "base_experience": 286,
	  "order": 336,
	  "is_default": false
	},
	{
	  "id": 10065,
	  "identifier": "sceptile-mega",
	  "species_id": 254,
	  "height": 19,
	  "weight": 552,
	  "base_experience": 284,
	  "order": 328,
	  "is_default": false
	},
	{
	  "id": 10066,
	  "identifier": "sableye-mega",
	  "species_id": 302,
	  "height": 5,
	  "weight": 1610,
	  "base_experience": 168,
	  "order": 382,
	  "is_default": false
	},
	{
	  "id": 10067,
	  "identifier": "altaria-mega",
	  "species_id": 334,
	  "height": 15,
	  "weight": 206,
	  "base_experience": 207,
	  "order": 423,
	  "is_default": false
	},
	{
	  "id": 10068,
	  "identifier": "gallade-mega",
	  "species_id": 475,
	  "height": 16,
	  "weight": 564,
	  "base_experience": 278,
	  "order": 361,
	  "is_default": false
	},
	{
	  "id": 10069,
	  "identifier": "audino-mega",
	  "species_id": 531,
	  "height": 15,
	  "weight": 320,
	  "base_experience": 425,
	  "order": 623,
	  "is_default": false
	},
	{
	  "id": 10070,
	  "identifier": "sharpedo-mega",
	  "species_id": 319,
	  "height": 25,
	  "weight": 1303,
	  "base_experience": 196,
	  "order": 406,
	  "is_default": false
	},
	{
	  "id": 10071,
	  "identifier": "slowbro-mega",
	  "species_id": 80,
	  "height": 20,
	  "weight": 1200,
	  "base_experience": 207,
	  "order": 116,
	  "is_default": false
	},
	{
	  "id": 10072,
	  "identifier": "steelix-mega",
	  "species_id": 208,
	  "height": 105,
	  "weight": 7400,
	  "base_experience": 214,
	  "order": 138,
	  "is_default": false
	},
	{
	  "id": 10073,
	  "identifier": "pidgeot-mega",
	  "species_id": 18,
	  "height": 22,
	  "weight": 505,
	  "base_experience": 261,
	  "order": 24,
	  "is_default": false
	},
	{
	  "id": 10074,
	  "identifier": "glalie-mega",
	  "species_id": 362,
	  "height": 21,
	  "weight": 3502,
	  "base_experience": 203,
	  "order": 458,
	  "is_default": false
	},
	{
	  "id": 10075,
	  "identifier": "diancie-mega",
	  "species_id": 719,
	  "height": 11,
	  "weight": 278,
	  "base_experience": 315,
	  "order": 834,
	  "is_default": false
	},
	{
	  "id": 10076,
	  "identifier": "metagross-mega",
	  "species_id": 376,
	  "height": 25,
	  "weight": 9429,
	  "base_experience": 315,
	  "order": 475,
	  "is_default": false
	},
	{
	  "id": 10077,
	  "identifier": "kyogre-primal",
	  "species_id": 382,
	  "height": 98,
	  "weight": 4300,
	  "base_experience": 347,
	  "order": 484,
	  "is_default": false
	},
	{
	  "id": 10078,
	  "identifier": "groudon-primal",
	  "species_id": 383,
	  "height": 50,
	  "weight": 9997,
	  "base_experience": 347,
	  "order": 486,
	  "is_default": false
	},
	{
	  "id": 10079,
	  "identifier": "rayquaza-mega",
	  "species_id": 384,
	  "height": 108,
	  "weight": 3920,
	  "base_experience": 351,
	  "order": 488,
	  "is_default": false
	},
	{
	  "id": 10080,
	  "identifier": "pikachu-rock-star",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 37,
	  "is_default": false
	},
	{
	  "id": 10081,
	  "identifier": "pikachu-belle",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 38,
	  "is_default": false
	},
	{
	  "id": 10082,
	  "identifier": "pikachu-pop-star",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 39,
	  "is_default": false
	},
	{
	  "id": 10083,
	  "identifier": "pikachu-phd",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 40,
	  "is_default": false
	},
	{
	  "id": 10084,
	  "identifier": "pikachu-libre",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 41,
	  "is_default": false
	},
	{
	  "id": 10085,
	  "identifier": "pikachu-cosplay",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 36,
	  "is_default": false
	},
	{
	  "id": 10086,
	  "identifier": "hoopa-unbound",
	  "species_id": 720,
	  "height": 65,
	  "weight": 4900,
	  "base_experience": 306,
	  "order": 836,
	  "is_default": false
	},
	{
	  "id": 10087,
	  "identifier": "camerupt-mega",
	  "species_id": 323,
	  "height": 25,
	  "weight": 3205,
	  "base_experience": 196,
	  "order": 411,
	  "is_default": false
	},
	{
	  "id": 10088,
	  "identifier": "lopunny-mega",
	  "species_id": 428,
	  "height": 13,
	  "weight": 283,
	  "base_experience": 203,
	  "order": 535,
	  "is_default": false
	},
	{
	  "id": 10089,
	  "identifier": "salamence-mega",
	  "species_id": 373,
	  "height": 18,
	  "weight": 1126,
	  "base_experience": 315,
	  "order": 471,
	  "is_default": false
	},
	{
	  "id": 10090,
	  "identifier": "beedrill-mega",
	  "species_id": 15,
	  "height": 14,
	  "weight": 405,
	  "base_experience": 223,
	  "order": 20,
	  "is_default": false
	},
	{
	  "id": 10091,
	  "identifier": "rattata-alola",
	  "species_id": 19,
	  "height": 3,
	  "weight": 38,
	  "base_experience": 51,
	  "order": 26,
	  "is_default": false
	},
	{
	  "id": 10092,
	  "identifier": "raticate-alola",
	  "species_id": 20,
	  "height": 7,
	  "weight": 255,
	  "base_experience": 145,
	  "order": 28,
	  "is_default": false
	},
	{
	  "id": 10093,
	  "identifier": "raticate-totem-alola",
	  "species_id": 20,
	  "height": 14,
	  "weight": 1050,
	  "base_experience": 145,
	  "order": 29,
	  "is_default": false
	},
	{
	  "id": 10094,
	  "identifier": "pikachu-original-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 36,
	  "is_default": false
	},
	{
	  "id": 10095,
	  "identifier": "pikachu-hoenn-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 37,
	  "is_default": false
	},
	{
	  "id": 10096,
	  "identifier": "pikachu-sinnoh-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 38,
	  "is_default": false
	},
	{
	  "id": 10097,
	  "identifier": "pikachu-unova-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 39,
	  "is_default": false
	},
	{
	  "id": 10098,
	  "identifier": "pikachu-kalos-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 40,
	  "is_default": false
	},
	{
	  "id": 10099,
	  "identifier": "pikachu-alola-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 41,
	  "is_default": false
	},
	{
	  "id": 10100,
	  "identifier": "raichu-alola",
	  "species_id": 26,
	  "height": 7,
	  "weight": 210,
	  "base_experience": 218,
	  "order": 44,
	  "is_default": false
	},
	{
	  "id": 10101,
	  "identifier": "sandshrew-alola",
	  "species_id": 27,
	  "height": 7,
	  "weight": 400,
	  "base_experience": 60,
	  "order": 46,
	  "is_default": false
	},
	{
	  "id": 10102,
	  "identifier": "sandslash-alola",
	  "species_id": 28,
	  "height": 12,
	  "weight": 550,
	  "base_experience": 158,
	  "order": 48,
	  "is_default": false
	},
	{
	  "id": 10103,
	  "identifier": "vulpix-alola",
	  "species_id": 37,
	  "height": 6,
	  "weight": 99,
	  "base_experience": 60,
	  "order": 59,
	  "is_default": false
	},
	{
	  "id": 10104,
	  "identifier": "ninetales-alola",
	  "species_id": 38,
	  "height": 11,
	  "weight": 199,
	  "base_experience": 177,
	  "order": 61,
	  "is_default": false
	},
	{
	  "id": 10105,
	  "identifier": "diglett-alola",
	  "species_id": 50,
	  "height": 2,
	  "weight": 10,
	  "base_experience": 53,
	  "order": 77,
	  "is_default": false
	},
	{
	  "id": 10106,
	  "identifier": "dugtrio-alola",
	  "species_id": 51,
	  "height": 7,
	  "weight": 666,
	  "base_experience": 149,
	  "order": 79,
	  "is_default": false
	},
	{
	  "id": 10107,
	  "identifier": "meowth-alola",
	  "species_id": 52,
	  "height": 4,
	  "weight": 42,
	  "base_experience": 58,
	  "order": 81,
	  "is_default": false
	},
	{
	  "id": 10108,
	  "identifier": "persian-alola",
	  "species_id": 53,
	  "height": 11,
	  "weight": 330,
	  "base_experience": 154,
	  "order": 83,
	  "is_default": false
	},
	{
	  "id": 10109,
	  "identifier": "geodude-alola",
	  "species_id": 74,
	  "height": 4,
	  "weight": 203,
	  "base_experience": 60,
	  "order": 107,
	  "is_default": false
	},
	{
	  "id": 10110,
	  "identifier": "graveler-alola",
	  "species_id": 75,
	  "height": 10,
	  "weight": 1100,
	  "base_experience": 137,
	  "order": 109,
	  "is_default": false
	},
	{
	  "id": 10111,
	  "identifier": "golem-alola",
	  "species_id": 76,
	  "height": 17,
	  "weight": 3160,
	  "base_experience": 223,
	  "order": 111,
	  "is_default": false
	},
	{
	  "id": 10112,
	  "identifier": "grimer-alola",
	  "species_id": 88,
	  "height": 7,
	  "weight": 420,
	  "base_experience": 65,
	  "order": 127,
	  "is_default": false
	},
	{
	  "id": 10113,
	  "identifier": "muk-alola",
	  "species_id": 89,
	  "height": 10,
	  "weight": 520,
	  "base_experience": 175,
	  "order": 129,
	  "is_default": false
	},
	{
	  "id": 10114,
	  "identifier": "exeggutor-alola",
	  "species_id": 103,
	  "height": 109,
	  "weight": 4156,
	  "base_experience": 186,
	  "order": 147,
	  "is_default": false
	},
	{
	  "id": 10115,
	  "identifier": "marowak-alola",
	  "species_id": 105,
	  "height": 10,
	  "weight": 340,
	  "base_experience": 149,
	  "order": 150,
	  "is_default": false
	},
	{
	  "id": 10116,
	  "identifier": "greninja-battle-bond",
	  "species_id": 658,
	  "height": 15,
	  "weight": 400,
	  "base_experience": 239,
	  "order": 760,
	  "is_default": false
	},
	{
	  "id": 10117,
	  "identifier": "greninja-ash",
	  "species_id": 658,
	  "height": 15,
	  "weight": 400,
	  "base_experience": 288,
	  "order": 761,
	  "is_default": false
	},
	{
	  "id": 10118,
	  "identifier": "zygarde-10",
	  "species_id": 718,
	  "height": 12,
	  "weight": 335,
	  "base_experience": 219,
	  "order": 830,
	  "is_default": false
	},
	{
	  "id": 10119,
	  "identifier": "zygarde-50",
	  "species_id": 718,
	  "height": 50,
	  "weight": 3050,
	  "base_experience": 270,
	  "order": 831,
	  "is_default": false
	},
	{
	  "id": 10120,
	  "identifier": "zygarde-complete",
	  "species_id": 718,
	  "height": 45,
	  "weight": 6100,
	  "base_experience": 319,
	  "order": 832,
	  "is_default": false
	},
	{
	  "id": 10121,
	  "identifier": "gumshoos-totem",
	  "species_id": 735,
	  "height": 14,
	  "weight": 600,
	  "base_experience": 146,
	  "order": 852,
	  "is_default": false
	},
	{
	  "id": 10122,
	  "identifier": "vikavolt-totem",
	  "species_id": 738,
	  "height": 26,
	  "weight": 1475,
	  "base_experience": 225,
	  "order": 856,
	  "is_default": false
	},
	{
	  "id": 10123,
	  "identifier": "oricorio-pom-pom",
	  "species_id": 741,
	  "height": 6,
	  "weight": 34,
	  "base_experience": 167,
	  "order": 860,
	  "is_default": false
	},
	{
	  "id": 10124,
	  "identifier": "oricorio-pau",
	  "species_id": 741,
	  "height": 6,
	  "weight": 34,
	  "base_experience": 167,
	  "order": 861,
	  "is_default": false
	},
	{
	  "id": 10125,
	  "identifier": "oricorio-sensu",
	  "species_id": 741,
	  "height": 6,
	  "weight": 34,
	  "base_experience": 167,
	  "order": 862,
	  "is_default": false
	},
	{
	  "id": 10126,
	  "identifier": "lycanroc-midnight",
	  "species_id": 745,
	  "height": 11,
	  "weight": 250,
	  "base_experience": 170,
	  "order": 869,
	  "is_default": false
	},
	{
	  "id": 10127,
	  "identifier": "wishiwashi-school",
	  "species_id": 746,
	  "height": 82,
	  "weight": 786,
	  "base_experience": 217,
	  "order": 872,
	  "is_default": false
	},
	{
	  "id": 10128,
	  "identifier": "lurantis-totem",
	  "species_id": 754,
	  "height": 15,
	  "weight": 580,
	  "base_experience": 168,
	  "order": 882,
	  "is_default": false
	},
	{
	  "id": 10129,
	  "identifier": "salazzle-totem",
	  "species_id": 758,
	  "height": 21,
	  "weight": 810,
	  "base_experience": 168,
	  "order": 887,
	  "is_default": false
	},
	{
	  "id": 10130,
	  "identifier": "minior-orange-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 904,
	  "is_default": false
	},
	{
	  "id": 10131,
	  "identifier": "minior-yellow-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 905,
	  "is_default": false
	},
	{
	  "id": 10132,
	  "identifier": "minior-green-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 906,
	  "is_default": false
	},
	{
	  "id": 10133,
	  "identifier": "minior-blue-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 907,
	  "is_default": false
	},
	{
	  "id": 10134,
	  "identifier": "minior-indigo-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 908,
	  "is_default": false
	},
	{
	  "id": 10135,
	  "identifier": "minior-violet-meteor",
	  "species_id": 774,
	  "height": 3,
	  "weight": 400,
	  "base_experience": 154,
	  "order": 909,
	  "is_default": false
	},
	{
	  "id": 10136,
	  "identifier": "minior-red",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 910,
	  "is_default": false
	},
	{
	  "id": 10137,
	  "identifier": "minior-orange",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 911,
	  "is_default": false
	},
	{
	  "id": 10138,
	  "identifier": "minior-yellow",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 912,
	  "is_default": false
	},
	{
	  "id": 10139,
	  "identifier": "minior-green",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 913,
	  "is_default": false
	},
	{
	  "id": 10140,
	  "identifier": "minior-blue",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 914,
	  "is_default": false
	},
	{
	  "id": 10141,
	  "identifier": "minior-indigo",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 915,
	  "is_default": false
	},
	{
	  "id": 10142,
	  "identifier": "minior-violet",
	  "species_id": 774,
	  "height": 3,
	  "weight": 3,
	  "base_experience": 175,
	  "order": 916,
	  "is_default": false
	},
	{
	  "id": 10143,
	  "identifier": "mimikyu-busted",
	  "species_id": 778,
	  "height": 2,
	  "weight": 7,
	  "base_experience": 167,
	  "order": 922,
	  "is_default": false
	},
	{
	  "id": 10144,
	  "identifier": "mimikyu-totem-disguised",
	  "species_id": 778,
	  "height": 4,
	  "weight": 28,
	  "base_experience": 167,
	  "order": 923,
	  "is_default": false
	},
	{
	  "id": 10145,
	  "identifier": "mimikyu-totem-busted",
	  "species_id": 778,
	  "height": 4,
	  "weight": 28,
	  "base_experience": 167,
	  "order": 924,
	  "is_default": false
	},
	{
	  "id": 10146,
	  "identifier": "kommo-o-totem",
	  "species_id": 784,
	  "height": 24,
	  "weight": 2075,
	  "base_experience": 270,
	  "order": 931,
	  "is_default": false
	},
	{
	  "id": 10147,
	  "identifier": "magearna-original",
	  "species_id": 801,
	  "height": 10,
	  "weight": 805,
	  "base_experience": 270,
	  "order": 952,
	  "is_default": false
	},
	{
	  "id": 10148,
	  "identifier": "pikachu-partner-cap",
	  "species_id": 25,
	  "height": 4,
	  "weight": 60,
	  "base_experience": 112,
	  "order": 42,
	  "is_default": false
	},
	{
	  "id": 10149,
	  "identifier": "marowak-totem",
	  "species_id": 105,
	  "height": 17,
	  "weight": 980,
	  "base_experience": 149,
	  "order": 151,
	  "is_default": false
	},
	{
	  "id": 10150,
	  "identifier": "ribombee-totem",
	  "species_id": 743,
	  "height": 4,
	  "weight": 20,
	  "base_experience": 162,
	  "order": 865,
	  "is_default": false
	},
	{
	  "id": 10151,
	  "identifier": "rockruff-own-tempo",
	  "species_id": 744,
	  "height": 5,
	  "weight": 92,
	  "base_experience": 56,
	  "order": 867,
	  "is_default": false
	},
	{
	  "id": 10152,
	  "identifier": "lycanroc-dusk",
	  "species_id": 745,
	  "height": 8,
	  "weight": 250,
	  "base_experience": 170,
	  "order": 870,
	  "is_default": false
	},
	{
	  "id": 10153,
	  "identifier": "araquanid-totem",
	  "species_id": 752,
	  "height": 31,
	  "weight": 2175,
	  "base_experience": 159,
	  "order": 879,
	  "is_default": false
	},
	{
	  "id": 10154,
	  "identifier": "togedemaru-totem",
	  "species_id": 777,
	  "height": 6,
	  "weight": 130,
	  "base_experience": 152,
	  "order": 920,
	  "is_default": false
	},
	{
	  "id": 10155,
	  "identifier": "necrozma-dusk",
	  "species_id": 800,
	  "height": 38,
	  "weight": 4600,
	  "base_experience": 306,
	  "order": 948,
	  "is_default": false
	},
	{
	  "id": 10156,
	  "identifier": "necrozma-dawn",
	  "species_id": 800,
	  "height": 42,
	  "weight": 3500,
	  "base_experience": 306,
	  "order": 949,
	  "is_default": false
	},
	{
	  "id": 10157,
	  "identifier": "necrozma-ultra",
	  "species_id": 800,
	  "height": 75,
	  "weight": 2300,
	  "base_experience": 339,
	  "order": 950,
	  "is_default": false
	}
]`

// AllPokemon returns a Pokedex filled with all the Pokemon
func AllPokemon() (*Pokedex, error) {
	var allPokemon []*Pokemon
	err := json.NewDecoder(strings.NewReader(PokedexJSON)).Decode(&allPokemon)
	if err != nil {
		return nil, err
	}
	dex := Pokedex{Pokemon: allPokemon}
	return &dex, nil
}
