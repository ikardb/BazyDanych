<script>
    import { onMount } from 'svelte';

    let stores = [
        { id: 1, name: "Sklep 1" },
        { id: 2, name: "Sklep 2" },
        { id: 3, name: "Sklep 3" },
    ];
    let selectedStoreId = null;
    let selectedOrderId = null;
    let newDiningMenuId = "";
    let newMenuPositionId = "";
    let newProductCount = null;
    let showOrders = false;
    let positionsAdding = false;
    let orders = [];
    let orderPositions = [];
    let diningMenus = [];
    let menuPositions = [];
    let error = null;
    let selectedOrder;

    $: selectedOrder = orders.find(order => order.id_zamowienia === selectedOrderId);
    
    const fetchOrdersByShopId  = async (storeId) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getOrdersByShopId/${storeId}`);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            orders = result.data;
        } catch (err) {
            error = err.message;
        }
    };

    const fetchOrderPositions = async (orderId) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getOrderPositions/${orderId}`);
            const result = await response.json();
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }
            orderPositions = result.data;
        } catch (err) {
            console.error(err.message);
        }
    };

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

    const fetchMenuPositions = async (id) => {
        try {
            const response = await fetch(`http://localhost:8080/api/getDiningMenuPositions/${id}`);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            menuPositions = result.data;
        } catch (err) {
            error = err.message;
        }
    };

    const migrateToStock = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/migrateToStock/${selectedOrderId}`, {
                method: 'POST',
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

        } catch (err) {
            error = err.message;
        }
    };

    const addOrder = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/createOrder', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id_sklepu: selectedStoreId,
                    id_uzytkownika: 1
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }
            fetchOrdersByShopId(selectedStoreId);
        } catch (err) {
            alert(err.message);
        }
    };

    const addOrderPosition = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/createOrderPosition', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id_zamowienia: selectedOrderId,
                    id_pozycji_jadlospisu: newMenuPositionId,
                    ilosc_produktu: newProductCount
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            fetchOrderPositions(selectedOrderId);
            fetchOrdersByShopId(selectedStoreId);
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
        showOrders = true;
        selectedStoreId = storeId;
        orders = [];
        selectedOrderId = null;
        orderPositions = [];
        await fetchOrdersByShopId(storeId);
    };

    const handleOrderClick = async (orderId) => {
        showOrders = false;
        selectedOrderId = orderId;
        orderPositions = [];
        await fetchOrderPositions(orderId);
    };

    const addingPosition = async () => {
        console.log('Dodawanie pozycji...');
        positionsAdding = true;
        await fetchDiningMenus();
    }

    const loadMenuPositions = async () => {
        await fetchMenuPositions(newDiningMenuId);
    }
</script>

<main>
    <h1>Zamówienia</h1>

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
        <div class="orders-container">
            <h2>Zamówienia dla Sklepu #{selectedStoreId}</h2>
            <button on:click={() => (selectedStoreId = null)}>Wróć do wyboru sklepu</button>
            {#if orders.length === 0}
                <p class="loading">Brak zamówień dla tego sklepu.</p>
                <button on:click={addOrder}>Dodaj zamówienie</button>
            {:else if showOrders}
                <button on:click={addOrder} style="margin-top: 10px;">Dodaj zamówienie</button>
                <ul>
                    {#each orders as order}
                        <li>
                            <button class="order-item" on:click={() => handleOrderClick(order.id_zamowienia)}>
                                <strong>ID Zamówienia:</strong> {order.id_zamowienia} <br>
                                <strong>Użytkownik:</strong> {order.id_uzytkownika} <br>
                                <strong>Data:</strong> {formatDate(order.data_zamowienia)} <br>
                                <strong>Kwota:</strong> {order.koszt_zamowienia} zł <br>
                                <strong>Czy wczytane do stanu:</strong> {order.wczytane_do_stanu}
                            </button>
                        </li>
                    {/each}
                </ul>
            {/if}

            {#if selectedOrderId}
                <div class="order-details">
                    <h3>Pozycje zamówienia #{selectedOrderId}</h3>
                    <button on:click={() => {selectedOrderId = null; showOrders = true; positionsAdding = false}}>Wróć do zamówień</button>
                    {#if !selectedOrder.wczytane_do_stanu && selectedOrder.koszt_zamowienia != 0}
                        <button on:click={migrateToStock}>Wczytaj zamówienie do stanu</button>
                    {/if}
                    {#if orderPositions.length === 0}
                        <p class="loading">Brak pozycji w tym zamówieniu.</p>
                        <button on:click={addingPosition}>Dodaj pozycję</button>
                        <div class="add-position">
                            {#if positionsAdding}
                                <select id="diningMenu" bind:value={newDiningMenuId} on:change={loadMenuPositions}>
                                    <option value="" disabled hidden selected>--Wybierz jadłospis--</option>
                                    {#each diningMenus as menu}
                                        <option value={menu.id_jadlospisu}>
                                            ID #{menu.id_jadlospisu}
                                        </option>
                                    {/each}
                                </select>

                                <select id="menuPosition" bind:value={newMenuPositionId}>
                                    <option value="" disabled selected>--Wybierz pozycję--</option>
                                    {#each menuPositions as position}
                                        <option value={position.id_pozycji_jadlospisu}>
                                            {position.nazwa}
                                        </option>
                                    {/each}
                                </select>
                    
                                <input type="number" min="1" bind:value={newProductCount} placeholder="Liczba produktów" />
                                <button on:click={addOrderPosition}>Dodaj</button>
                                <button on:click={() => {positionsAdding = false}}>Anuluj</button>
                            {/if}
                        </div>
                    {:else}
                        <ul>
                            {#each orderPositions as position}
                                <li class="position-item">
                                    <strong>ID pozycji:</strong> {position.id_pozycji_zamowienia} <br>
                                    <strong>Produkt:</strong> {position.nazwa} <br>
                                    <strong>Ilość:</strong> {position.ilosc_produktu} <br>
                                    <strong>Cena za sztukę:</strong> {position.cena} zł <br>
                                </li>
                            {/each}
                        </ul>
                        <button on:click={addingPosition}>Dodaj pozycję</button>
                        <div class="add-position">
                            {#if positionsAdding}
                                <select id="diningMenu" bind:value={newDiningMenuId} on:change={loadMenuPositions}>
                                    <option value="" disabled hidden selected>--Wybierz jadłospis--</option>
                                    {#each diningMenus as menu}
                                        <option value={menu.id_jadlospisu}>
                                            ID #{menu.id_jadlospisu}
                                        </option>
                                    {/each}
                                </select>

                                <select id="menuPosition" bind:value={newMenuPositionId}>
                                    <option value="" disabled selected>--Wybierz pozycję--</option>
                                    {#each menuPositions as position}
                                        <option value={position.id_pozycji_jadlospisu}>
                                            {position.nazwa}
                                        </option>
                                    {/each}
                                </select>
                    
                                <input type="number" min="1" bind:value={newProductCount} placeholder="Liczba produktów" />
                                <button on:click={addOrderPosition}>Dodaj</button>
                                <button on:click={() => {positionsAdding = false}}>Anuluj</button>
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

    .orders-container {
        margin-top: 2em;
        max-width: 800px;
        margin-left: auto;
        margin-right: auto;
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .orders-container ul {
        list-style-type: none;
        display: flex;
        flex-direction: column-reverse;
        width: 100%;
        padding: 0;
        margin: 0;
    }


    .order-item {
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

    .order-item:hover {
        background: #eaeaea;
    }

    .order-details {
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
