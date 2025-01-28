# IkarDB
This is a repository for a university project in which we have to develop a database
for the food-service company (locally cooked & distributed regional cousine in self-managed shops).

# Transactions' descriptions
<details>
	<summary>User</summary>

Args:	imie,
		nazwisko,
		admin,
		login,
		haslo,
		id_sklepu

#### Transaction process

###### Creation
	1. Creates a user with given args	

###### Authentication
	1. Checks if there is a user in the database with given credentials
	2. If yes: go ahead. If not: ask again (stop for 5 mins after 5 attempts - imo ok)

</details>

<details>
	<summary>Order</summary>


#### Transaction process


1. Order creation
2. Create and add *order positions* (from the dining menu position table)
3. Calculate worth of *order positions* and write it to the order


LATER irl:
__CONFIRMATION THE ORDER AT ARRIVAL__


</details>

<details>
	<summary>Dining menu</summary>

#### Transaction process

1. Dining menu creation
2. Create and add *dining menu* positions (from the products table)

</details>

<details>
	<summary>Sale</summary>

#### Transaction process

1. Sale creation
2. Create and add *sale positions* (from the products table)
3. Calculate worth of a sale and write it accordingly

</details>

<details>
	<summary>Stock level check</summary>
	
Args:	id_sklepu
		id_produktu
		ilosc

#### Transaction process

1. Stock level check creation

Desc.:	must happen __every time__ at __sale creation__ and __order confirmation__ with current implementation 

Possible improvements: adding a date field to a table as a primary key

</details>


<details>
	<summary>Revenue from a timestep</summary>

#### Transaction process

1. Track every sale within the given timestep (default can be month idk)
2. If not asked for a specific store, sum up the sale worth
3. Else from previously tracked sales track those that happened in the desired store
4. Sum them up

</details>

<details>
	<summary>Third-party company</summary>

#### Transaction process

1. Create a company
2. Add an offer

</details>

<details>
	<summary>Third-party company offer position</summary>

#### Transaction process

1. Create products that a company has to offer
2. Create the offer

</details>
