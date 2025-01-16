<script>
    import { onMount } from 'svelte';

    let diningMenus = [];
    let error = null;

    // Funkcja do pobierania danych z API
    const fetchDiningMenus = async () => {
    try {
        const response = await fetch('http://localhost:8080/api/diningMenus'); // lub 'http://localhost:8080/api/diningMenus', jeśli proxy nie działa

        if (!response.ok) {
        throw new Error('Błąd pobierania danych');
        }

        const result = await response.json(); // Pobranie odpowiedzi
        diningMenus = result.data; // Przypisanie pola `data` do `diningMenus`
    } catch (err) {
        error = err.message;
    }
};

    // Użycie onMount do wywołania fetch po załadowaniu komponentu
    onMount(() => {
      fetchDiningMenus();
    });
</script>

<main>
    <h1>Jadłospis</h1>
  
    {#if error}
      <p class="error">{error}</p>
    {:else if !diningMenus || diningMenus.length === 0}
      <p class="loading">Brak dostępnych jadłospisów.</p>
    {:else}
      <div class="menu-container">
        {#each diningMenus as menu}
          <div class="menu-item">
            <h2>Jadłospis #{menu.id_jadlospisu}</h2>
            <p><strong>ID Kuchni:</strong> {menu.id_kuchni}</p>
            <p><strong>ID Firmy:</strong> {menu.id_firmy || 'Brak'}</p>
          </div>
        {/each}
      </div>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        font-family: Arial, sans-serif;
    }

    h1 {
        font-size: 2.5em;
        color: #333;
        margin-bottom: 1em;
    }

    .error {
        color: red;
        font-size: 1.2em;
    }

    .loading {
        color: #555;
        font-size: 1.2em;
    }

    .menu-container {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1.5em;
        margin-top: 1em;
    }

    .menu-item {
        border: 1px solid #ddd;
        border-radius: 8px;
        padding: 1em;
        background-color: #f9f9f9;
        width: 250px;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        text-align: left;
    }

    .menu-item h2 {
        font-size: 1.5em;
        margin-bottom: 0.5em;
        color: #444;
    }

    .menu-item p {
        margin: 0.3em 0;
        font-size: 1em;
        color: #666;
    }

    .menu-item strong {
        color: #333;
    }
</style>
  
  