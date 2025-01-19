<script>
    let stores = [
        { id: 1, name: "Sklep 1" },
        { id: 2, name: "Sklep 2" },
        { id: 3, name: "Sklep 3" },
    ];

    let selectedStoreId = null;
    let selectedSaleId = null;
    let showSales = false;
    let error = null;
    let sales = [];
    let salePositions = [];
    let positionsAdding = false;

    //dodawanie produktow do sprzedazy
    let availableProducts = []; 
    let selectedProductId = "";
    let selectedQuantity = 1;
    let maxQuantity = 1;

    const fetchSalesByShopId = async (storeId) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getSalesByShopId/${storeId}`);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            sales = result.data;
        } catch (err) {
            error = err.message;
        }
    };

    const fetchSalePositions = async (saleId) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getSalePositions/${saleId}`);
            const result = await response.json();
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }
            salePositions = result.data;
        } catch (err) {
            console.error(err.message);
        }
    };

    const fetchAvailableProducts = async (storeId) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getStockLevelByShopID/${storeId}`);
            const result = await response.json();
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }
            availableProducts = result.data;
        } catch (err) {
            console.error(err.message);
        }
    };

    function handleProductChange(event) {
        const selectedId = +event.target.value;
        selectedProductId = selectedId;

        const selectedProduct = availableProducts.find(
            (product) => product.id_produktu === selectedId
        );

        if (selectedProduct) {
            maxQuantity = selectedProduct.ilosc;
            selectedQuantity = Math.min(selectedQuantity, maxQuantity);
        }
    }

    const addPositionToSale = async () => {
        if (!selectedProductId || selectedQuantity < 1) {
            alert("Wybierz produkt i ilość.");
            return;
        }
        try {
            const response = await fetch("http://localhost:8080/api/createSalePosition", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    id_sprzedazy: selectedSaleId,
                    id_produktu: selectedProductId,
                    ilosc: selectedQuantity
                }),
            });
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            fetchAvailableProducts(selectedStoreId);
            fetchSalePositions(selectedSaleId);
            selectedProductId = "";
            selectedQuantity = 1;
        } catch (err) {
            console.error(err.message);
            alert(err.message);
        }
    };

    const cancelAdding = () => {
        positionsAdding = false;
        selectedProductId = "";
        selectedQuantity = 1;
        fetchSalesByShopId(selectedStoreId);
    };

    const addSale = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/createSale', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id_sklepu: selectedStoreId,
                    id_uzytkownika: 1,
                    kwota_transakcji: 0.0,
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }
            fetchSalesByShopId(selectedStoreId);
        } catch (err) {
            alert(err.message);
        }
    };

    const formatDate = (isoDate) => {
        const date = new Date(isoDate);
        return date.toLocaleString("pl-PL", {
            year: "numeric",
            month: "long",
            day: "numeric",
            hour: "2-digit",
            minute: "2-digit",
        });
    };

    const handleStoreClick = async (storeId) => {
        showSales = true;
        selectedStoreId = storeId;
        sales = [];
        selectedSaleId = null;
        salePositions = [];
        await fetchSalesByShopId(selectedStoreId);
    };

    const handleSaleClick = async (saleId) => {
        showSales = false;
        selectedSaleId = saleId;
        salePositions = [];
        await fetchSalePositions(selectedSaleId);
    };

    const addingPosition = async () => {
        positionsAdding = true;
        fetchAvailableProducts(selectedStoreId);
    }

</script>

<main>
    <h1>Sprzedaże</h1>
    {#if error}
        <p class="error">{error}</p>
    {:else if !selectedStoreId}
        <div class="store-container">
            {#each stores as store}
                <button class="store-item" on:click={() => handleStoreClick(store.id)}>
                    <h2>{store.name}</h2>
                </button>
            {/each}
        </div>
    {:else}
    <div class="sales-container">
        <h2>Sprzedaże dla Sklepu #{selectedStoreId}</h2>
        <button on:click={() => (selectedStoreId = null)}>Wróć do wyboru sklepu</button>
        {#if sales.length === 0}
            <p class="loading">Brak sprzedaży dla tego sklepu.</p>
            <button on:click={addSale}>Dodaj sprzedaż</button>
        {:else if showSales}
            <button on:click={addSale} style="margin-top: 10px;">Dodaj sprzedaż</button>
            <ul>
                {#each sales as sale}
                    <li>
                        <button class="sale-item" on:click={() => handleSaleClick(sale.id_sprzedazy)}>
                            <strong>ID sprzedaży:</strong> {sale.id_sprzedazy} <br>
                            <strong>Użytkownik:</strong> {sale.id_uzytkownika} <br>
                            <strong>Data:</strong> {formatDate(sale.data_sprzedazy)} <br>
                            <strong>Kwota:</strong> {sale.kwota_transakcji} zł <br>
                        </button>
                    </li>
                {/each}
            </ul>
        {/if}

        {#if selectedSaleId}
        <div class="sale-details">
            <h3>Pozycje zamówienia #{selectedSaleId}</h3>
            <button on:click={() => {selectedSaleId = null; showSales = true; positionsAdding = false; fetchSalesByShopId(selectedStoreId)}}>Wróć do sprzedaży</button>
            {#if salePositions.length === 0}
                <p class="loading">Brak pozycji w tej sprzedaży.</p>
                <button on:click={addingPosition}>Dodaj pozycję</button>
                <div class="add-position">
                    {#if positionsAdding}
                        <select id="productSelect" bind:value={selectedProductId} on:change={handleProductChange}>
                            <option value="" disabled hidden selected>--Wybierz produkt--</option>
                            {#each availableProducts.filter(product => product.ilosc > 0) as product}
                                <option value={product.id_produktu}>
                                    {product.nazwa} (Dostępne: {product.ilosc})
                                </option>
                            {/each}
                        </select>
                    
                        {#if selectedProductId}
                            <div>
                                <label for="quantitySelect">Wybierz ilość:</label>
                                <input
                                    id="quantitySelect"
                                    type="number"
                                    min="1"
                                    max={maxQuantity}
                                    bind:value={selectedQuantity}
                                />
                            </div>
                        {/if}
            
                        <button on:click={addPositionToSale}>Dodaj</button>
                        <button on:click={cancelAdding}>Anuluj</button>
                    {/if}
                </div>
            {:else}
                <ul>
                    {#each salePositions as position}
                        <li class="position-item">
                            <strong>ID pozycji:</strong> {position.id_pozycji} <br>
                            <strong>Produkt:</strong> {position.nazwa} <br> 
                            <strong>Ilość:</strong> {position.ilosc} <br>
                            <strong>Cena jednostkowa:</strong> {position.cena_jednostkowa} zł <br>
                        </li>
                    {/each}
                </ul>
                <button on:click={addingPosition}>Dodaj pozycję</button>
                <div class="add-position">
                    {#if positionsAdding}
                        <select id="productSelect" bind:value={selectedProductId} on:change={handleProductChange}>
                            <option value="" disabled hidden selected>--Wybierz produkt--</option>
                            {#each availableProducts.filter(product => product.ilosc > 0) as product}
                                <option value={product.id_produktu}>
                                    {product.nazwa} (Dostępne: {product.ilosc})
                                </option>
                            {/each}
                        </select>
                    
                        {#if selectedProductId}
                            <div>
                                <label for="quantitySelect">Wybierz ilość:</label>
                                <input
                                    id="quantitySelect"
                                    type="number"
                                    min="1"
                                    max={maxQuantity}
                                    bind:value={selectedQuantity}
                                />
                            </div>
                        {/if}
            
                        <button on:click={addPositionToSale}>Dodaj</button>
                        <button on:click={cancelAdding}>Anuluj</button>
                    {/if}
                </div>
            {/if}
        </div>
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

    .sales-container {
        margin-top: 2em;
        max-width: 800px;
        margin-left: auto;
        margin-right: auto;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .sales-container ul {
        list-style-type: none;
        display: flex;
        flex-direction: column-reverse;
        width: 100%;
        padding: 0;
        margin: 0;
    }

    .sale-item {
        background: #f4f4f4;
        border: none;
        margin: 0.5em 0;
        padding: 1em;
        border-radius: 5px;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        cursor: pointer;
        text-align: left;
        width: 100%;
    }

    .sale-item:hover {
        background: #eaeaea;
    }

    .sale-details {
        margin-top: 1em;
        width: 100%;
        background: #fff;
        padding: 1em;
        border: 1px solid #ddd;
        border-radius: 5px;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .position-item {
        background: #f9f9f9;
        margin: 0.5em 0;
        padding: 0.5em;
        border-radius: 5px;
        border: 1px solid #ddd;
        text-align: left;
    }

    .add-position {
        display: flex;
        flex-direction: column;
        width: 200px;
    }
</style>
