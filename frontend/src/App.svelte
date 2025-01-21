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
		<Link to="/">Home</Link>
		<Link to="/diningMenu">Jadłospis</Link>
		<Link to="/orders">Zamówienia</Link>
		<Link to="/sales">Sprzedaże</Link>
		<Link to="/stockLevels">Stany magazynowe</Link>
		<Link to="/revenue">Utarg</Link>
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
}
</style>
  