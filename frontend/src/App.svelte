<script>
	import { Router, Route, Link } from 'svelte-routing';
	import Home from './routes/Home.svelte';
	import DiningMenu from './routes/DiningMenu.svelte';
	import Orders from './routes/Orders.svelte';
</script>

<Router>
	<nav>
		<Link to="/">Home</Link>
		<Link to="/diningMenu">Jadłospis</Link>
		<Link to="/orders">Zamówienia</Link>
	</nav>

	<Route path="/" component={Home} />
	<Route path="/diningMenu" component={DiningMenu} />
	<Route path="/orders" component={Orders} />
</Router>

<style>
nav {
	margin: 20px;
}
</style>
  