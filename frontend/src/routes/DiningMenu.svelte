<script>
    import { onMount } from 'svelte';

    let diningMenus = [];
    let error = null;
    let selectedMenuId = null;
    let menuPositions = [];
    let addingMenu = false;
    let addingPositions = false;
    let addingProduct = false;
    let companies = [];
    let kitchens = [];
    let products = [];
    let newProductId = "";
    let newMenuCompanyId = "";
    let newMenuKitchenId = "";
    let newProductName = "";
    let newProductPrice = null;

    // Funkcja do pobierania wszystkich jadłospisów
    const fetchDiningMenus = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/diningMenus');

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            diningMenus = result.data;
        } catch (err) {
            error = err.message;
        }
    };

    const fetchCompanies = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/thirdPartyCompanies');
            const result = await response.json();
            if (!response.ok) {
                throw new Error(result.message);
            }
            companies = result.data;
        } catch (err) {
            console.error('Błąd:', err);
            error = err.message;
        }
    };

    const fetchKitchens = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/kitchens');
            const result = await response.json();
            if (!response.ok) {
                throw new Error(result.message);
            }
            kitchens = result.data;
        } catch (err) {
            console.error('Błąd:', err);
            error = err.message;
        }
    };

    const fetchProducts = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/products');
            const result = await response.json();
            if (!response.ok) {
                throw new Error(result.message);
            }
            products = result.data;
        } catch (err) {
            console.error('Błąd:', err);
            error = err.message;
        }
    };

    // Funkcja do dodawania jadłospisu
    const addDiningMenu = async () => {
        if (!newMenuKitchenId && !newMenuCompanyId) {
            alert("Proszę wybrać id kuchni lub id firmy.");
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/api/createDiningMenu', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id_kuchni: newMenuKitchenId,
                    id_firmy: newMenuCompanyId
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            fetchDiningMenus();
            newMenuKitchenId = null;
            newMenuCompanyId = null;
            addingMenu = false;
        } catch (err) {
            alert(err.message);
        }
    };

    const addMenuPosition = async () => {
        if (!newProductId) {
            alert("Proszę wybrać produkt.");
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/api/createDiningMenuPosition', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id_jadlospisu: selectedMenuId,
                    id_produktu: newProductId
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            await fetchMenuPositions(selectedMenuId);
            newProductId = null;
            addingPositions = false;
            addingProduct = false;
        } catch (err) {
            alert(err.message);
        }
    };

    const addProduct = async () => {
        if (!newProductName || !newProductPrice) {
            alert("Proszę podać zarówno nazwę, jak i cenę produktu.");
        return;
    }
        try {
            const response = await fetch('http://localhost:8080/api/createProduct', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    nazwa: newProductName,
                    cena: newProductPrice
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            await fetchProducts();
            newProductId = null;
            addingProduct = false;
        } catch (err) {
            alert(err.message);
        }
    };

    const fetchMenuPositions = async (id) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getDiningMenuPositions/${id}`);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            menuPositions = result.data;
            addingPositions = false;
            addingProduct = false;
        } catch (err) {
            error = err.message;
        }
    };

    const menuAdding = () => {
        addingMenu = true;
        selectedMenuId = null;
        addingPositions = false;
        addingProduct = false;
    };

    const positionsAdding = () => {
        addingPositions = true;
    }

    const productAdding = () => {
        addingProduct = true;
    }

    onMount(() => {
        fetchDiningMenus();
        fetchCompanies();
        fetchKitchens();
        fetchProducts();
    });

    const handleMenuClick = async (id) => {
        selectedMenuId = id;
        addingMenu = false;
        menuPositions = [];
        await fetchMenuPositions(id);
    };
</script>

<main>
    <h1>Jadłospisy</h1>
    {#if error}
        <p class="error">{error}</p>
    {:else if !diningMenus || diningMenus.length === 0}
        <p class="loading">Brak dostępnych jadłospisów.</p>
    {:else}
        <div class="menu-container">
            {#each diningMenus as menu}
                <button class="menu-item" on:click={() => handleMenuClick(menu.id_jadlospisu)}>
                    <h2>Jadłospis #{menu.id_jadlospisu}</h2>
                    <p><strong>ID Kuchni:</strong> {menu.id_kuchni || 'Brak'}</p>
                    <p><strong>ID Firmy:</strong> {menu.id_firmy || 'Brak'}</p>
                </button>
            {/each}
        </div>

        <button class="add_menu_button" on:click={() => menuAdding()}>Dodaj jadłospis</button>

        {#if addingMenu}
            <div class="add-menu">
                <h2>Dodaj nowy jadłospis</h2>
                <label for="kitchen">Wybierz kuchnię</label>
                <select id="kitchen" bind:value={newMenuKitchenId} on:change={() => newMenuCompanyId = null}>
                    {#each kitchens as kitchen}
                        <option value={kitchen.id_kuchni}>
                            {kitchen.id_kuchni} - {kitchen.ulica}
                        </option>
                    {/each}
                </select>

                <label for="company">Wybierz firmę</label>
                <select id="company" bind:value={newMenuCompanyId} on:change={() => newMenuKitchenId = null}>
                    {#each companies as company}
                        <option value={company.id_firmy}>
                            {company.id_firmy} - {company.nazwa}
                        </option>
                    {/each}
                </select>
        
                <button on:click={addDiningMenu}>Dodaj</button>
            </div>
        {/if}

        {#if selectedMenuId}
            <div class="menu-details">
                <h2>Pozycje dla jadłospisu #{selectedMenuId}</h2>
                <button class="add_menu_button" on:click={() => positionsAdding()}>Dodaj pozycje</button>
                {#if addingPositions}
                    <div class="positionsHandlers">
                        <div class="add-positions">
                            <h2>Dodaj nową pozycję</h2>
                            <select id="position" bind:value={newProductId}>
                                <option value="" disabled hidden selected>--Wybierz produkt--</option>
                                {#each products as product}
                                    <option value={product.id_produktu}>
                                        {product.id_produktu} - {product.nazwa}
                                    </option>
                                {/each}
                            </select>
                            <button on:click={addMenuPosition}>Dodaj</button>
                            <button on:click={() => productAdding()}>Dodaj nowy produkt</button>
                            <button on:click={() => {addingPositions = false; addingProduct = false}}>Anuluj</button>
                        </div>
                        {#if addingProduct}
                            <div class="add-products">
                                <h2>Dodaj nowy produkt</h2>
                                <input id="productName" type="text" placeholder="Wpisz nazwę produktu" bind:value={newProductName}/>
                                
                                <input id="productPrice" type="number" placeholder="Wpisz cenę produktu" bind:value={newProductPrice} min="0" step="0.01"/>
                                <button on:click={addProduct}>Dodaj</button>
                                <button on:click={() => {addingProduct = false}}>Anuluj</button>
                            </div>
                        {/if}
                    </div>
                {/if}
                {#if menuPositions.length === 0}
                    <p class="loading">Jadłospis jest pusty</p>
                {:else}
                    <ul>
                        {#each menuPositions as position}
                            <li>
                                <strong>ID pozycji #</strong>{position.id_pozycji_jadlospisu} - {position.nazwa}
                            </li>
                        {/each}
                    </ul>
                {/if}
            </div>
        {/if}
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
        cursor: pointer;
        transition: transform 0.2s;
    }

    .menu-item:hover {
        transform: scale(1.05);
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

    .menu-details {
        margin-top: 2em;
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: left;
        max-width: 600px;
        margin-left: auto;
        margin-right: auto;
    }

    .menu-details h2 {
        color: #333;
        font-size: 1.8em;
        margin-bottom: 0;
    }

    .menu-details ul {
        list-style-type: none;
        padding: 0;
    }

    .menu-details li {
        background: #f4f4f4;
        margin: 0.5em 0;
        padding: 0.8em;
        border-radius: 5px;
        box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
    }

    .positionsHandlers {
        position: relative;
        display: flex;
        justify-content: center;
    }

    .add-positions {
        position: inherit;
        display: flex;
        width: 300px;
        flex-direction: column;
        justify-content: center;
        text-align: center;
    }

    .add-products {
        text-align: center;
        position: absolute;
        top: 0;
        left: 300px;
        width: 300px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin-left: 20px;
    }
</style>
