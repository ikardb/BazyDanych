<script>
	import { Router, Route, Link } from 'svelte-routing';
	import Home from './routes/Home.svelte';
	import DiningMenu from './routes/DiningMenu.svelte';
	import Orders from './routes/Orders.svelte';
	import Sales from './routes/Sales.svelte';
	import StockLevels from './routes/StockLevels.svelte';
	import Revenue from './routes/Revenue.svelte';
</script>

<Router>
	<nav>
		<a href='/' class='link' use:Link>Home</a>
		<a href='/diningMenu' class='link' use:Link>Jadłospis</a>
		<a href='/orders' class='link' use:Link>Zamówienia</a>
		<a href='/sales' class='link' use:Link>Sprzedaże</a>
		<a href='/stockLevels' class='link' use:Link>Stany magazynowe</a>
		<a href='/revenue' class='link' use:Link>Utarg</a>
	</nav>

	<Route path="/" component={Home} />
	<Route path="/diningMenu" component={DiningMenu} />
	<Route path="/orders" component={Orders} />
	<Route path="/sales" component={Sales} />
	<Route path="/stockLevels" component={StockLevels} />
	<Route path="/revenue" component={Revenue} />
</Router>

<style>
	nav {
		margin: 20px;
		padding: 15px;
		justify-content: center;
		display: flex;
		gap: 20px;
	}

	.link {
		display: block; /* Linki są blokami, co sprawia, że zachowują się jak przyciski */
		padding: 10px 20px; /* Wewnętrzne marginesy, aby linki wyglądały na większe */
		background-color: #42312e; /* Tło linków */
		color: white; /* Kolor tekstu */
		text-align: center; /* Wyśrodkowanie tekstu w linku */
		text-decoration: none; /* Usunięcie podkreślenia */
		border-radius: 5px; /* Zaokrąglenie rogów */
		font-size: 16px; /* Ustawienie rozmiaru czcionki */
		font-weight: bold; /* Pogrubienie tekstu */
		transition: background-color 0.3s, transform 0.2s; /* Płynne przejścia */
	}

  .link:hover {
    background-color: #615958;
    transform: translateY(-3px);
  }

  .link:active {
    transform: translateY(1px);
  }
</style>
  