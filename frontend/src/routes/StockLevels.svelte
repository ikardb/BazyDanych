<script>
    let stores = [
        { id: 1, name: "Sklep 1" },
        { id: 2, name: "Sklep 2" },
        { id: 3, name: "Sklep 3" },
    ];
    let error = null;
    let selectedStoreId = null;
    let showStockLevels = false;
    let stockLevels = [];
    let searchQuery = "";
    
    const fetchStockLevelsByShopId = async (storeId, query = "") => {
        try {
            const url = `http://localhost:8080/api/getStockLevelByShopID/${storeId}?q=${encodeURIComponent(query)}`;
            const response = await fetch(url);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            stockLevels = result.data;
            error = null;
        } catch (err) {
            error = err.message;
        }
    };

    const handleStoreClick = async (storeId) => {
        selectedStoreId = storeId;
        showStockLevels = true;
        searchQuery = "";
        await fetchStockLevelsByShopId(selectedStoreId);
    };

    const handleSearch = async () => {
        if (selectedStoreId) {
            await fetchStockLevelsByShopId(selectedStoreId, searchQuery);
        }
    };
</script>

<main>
    {#if error}
        <p class="error">{error}</p>
    {:else if !selectedStoreId}
        <h1>Stany magazynowe</h1>
        <div class="store-container">
            {#each stores as store}
                <button class="store-item" on:click={() => handleStoreClick(store.id)}>
                    <h2>{store.name}</h2>
                </button>
            {/each}
        </div>
    {:else}
        <div class="stocks-container">
            <h1>Stan dla sklepu #{selectedStoreId}</h1>
            <button on:click={() => (selectedStoreId = null)}>Wróć do wyboru sklepu</button>

            <input type="text" placeholder="Wyszukaj:" bind:value={searchQuery} on:input={handleSearch}/>
            {#if stockLevels.length === 0}
                <p class="loading">Brak produktów dla tego sklepu.</p>
            {:else if showStockLevels}
                <ul>
                    {#each stockLevels as stock}
                        <li>
                            <div class="stock-item">
                                <strong>Nazwa:</strong> {stock.nazwa} <br>
                                <strong>Ilość:</strong> {stock.ilosc} <br>
                            </div>
                        </li>
                    {/each}
                </ul>
            {/if}
        </div>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        font-family: Arial, sans-serif;
    }

    .error {
        color: red;
        font-size: 1.2em;
    }

    .loading {
        color: #555;
        font-size: 1.2em;
    }

    .store-container {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1.5em;
        margin-top: 1em;
    }

    .store-item {
        border: 1px solid #ddd;
        border-radius: 8px;
        padding: 1em;
        background-color: #f9f9f9;
        width: 250px;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        text-align: center;
        cursor: pointer;
        transition: transform 0.2s;
    }

    .store-item:hover {
        transform: scale(1.05);
    }

    .stocks-container {
        margin-top: 2em;
        max-width: 800px;
        margin-left: auto;
        margin-right: auto;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .stocks-container ul {
        list-style-type: none;
        display: flex;
        flex-direction: column-reverse;
        width: 100%;
        padding: 0;
        margin: 0;
    }

    .stock-item {
        background: #f4f4f4;
        border: none;
        margin: 0.5em 0;
        padding: 1em;
        border-radius: 5px;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        text-align: left;
        width: 100%;
    }

    input[type="text"] {
        padding: 0.5em;
        font-size: 1em;
        margin-bottom: 1em;
        border: 1px solid #ddd;
        border-radius: 5px;
        width: 80%;
        max-width: 400px;
    }

    input[type="text"]:focus {
        outline: none;
        border-color: #007bff;
    }
</style>