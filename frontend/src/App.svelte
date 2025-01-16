<script>
	import { Router, Route, Link } from 'svelte-routing';
	import Home from './routes/Home.svelte';
	import DiningMenu from './components/diningMenu.svelte';
</script>

<Router>
<nav>
	<Link to="/">Home</Link>
	<Link to="/diningMenu">Jadłospis</Link>
</nav>

<Route path="/" component={Home} />
<Route path="/diningMenu" component={DiningMenu} />
</Router>

<style>
nav {
	margin: 20px;
}
</style>
  