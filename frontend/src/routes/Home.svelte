<script>
    import { Link } from 'svelte-routing';
</script>

<main>
    <h1>IkarDB</h1>
    <p>Przejdź do <a href="/diningMenu">Jadłospisów</a>.</p>
    <p>Przejdź do <a href="/orders">Zamówień</a>.</p>
    <p>Przejdź do <a href="/sales">Sprzedaży</a>.</p>
    <p>Przejdź do <a href="/stockLevels">Stanów magazynowych</a>.</p>
    <p>Przejdź do <a href="/revenue">Podsumowania utargu</a>.</p>
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
  