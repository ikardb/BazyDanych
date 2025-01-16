<script>
    import { Link } from 'svelte-routing';
</script>

<main>
<h1>Witaj w mojej aplikacji!</h1>
<p>Przejdź do <Link to="/diningMenu">Jadłospisu</Link>.</p>
</main>

<style>
main {
    text-align: center;
}

h1 {
    font-size: 2em;
    color: #333;
}
</style>
  