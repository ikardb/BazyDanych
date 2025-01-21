<script>
    import { onMount } from "svelte";

    let error = null;

    let stores = [
        { id: 1, name: "Sklep 1" },
        { id: 2, name: "Sklep 2" },
        { id: 3, name: "Sklep 3" },
    ];
    let users = [];

    let selectedStoreId = "";
    let selectedUserId = "";
    let startDate = "";
    let endDate = "";

    let salesData = null;
    let revenue = null;

    const fetchUsers = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/Users`);

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            users = result.data;
        } catch (err) {
            error = err.message;
        }
    };

    const fetchSalesData = async (event) => {
        event.preventDefault();

        const startDateTime = startDate ? new Date(`${startDate}T00:00:00`).toISOString() : null;
        const endDateTime = endDate ? new Date(`${endDate}T23:59:59`).toISOString() : null;

        const requestData = {
            od_kiedy: startDateTime,
            do_kiedy: endDateTime,
            id_sklepu: selectedStoreId || null,
            id_uzytkownika: selectedUserId || null,
        };

        try {
            const response = await fetch("http://localhost:8080/api/sumSalesFromGivenTime", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(requestData),
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message);
            }

            const result = await response.json();
            salesData = result.data;
            revenue = result.utarg;
            error = null;
        } catch (err) {
            salesData = null;
            error = err.message;
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

    onMount(() => {
        fetchUsers();
    });
</script>

<main>
    <h1>Podsumowanie utargu</h1>

    <form on:submit={fetchSalesData}>
        <div>
            <label for="store">Wybierz sklep:</label>
            <select id="store" bind:value={selectedStoreId}>
                <option value="">Wszystkie sklepy</option>
                {#each stores as store}
                    <option value={store.id}>{store.name}</option>
                {/each}
            </select>
        </div>

        <div>
            <label for="user">Wybierz użytkownika:</label>
            <select id="user" bind:value={selectedUserId}>
                <option value="">Wszyscy użytkownicy</option>
                {#each users as user}
                    <option value={user.id_uzytkownika}>{user.imie} {user.nazwisko}</option>
                {/each}
            </select>
        </div>

        <div>
            <label for="start-date">Data od:</label>
            <input type="date" id="start-date" bind:value={startDate} />
        </div>

        <div>
            <label for="end-date">Data do:</label>
            <input type="date" id="end-date" bind:value={endDate} />
        </div>

        <button type="submit">Pobierz utarg</button>
    </form>

    {#if error}
        <p class="error">Błąd: {error}</p>
    {:else if salesData}
        <div class="result">
            <h2>Utarg: {revenue} zł</h2>
            <ul>
                {#each salesData as data}
                    <li>
                        <div class="sale-item">
                            <strong>ID sprzedaży:</strong> {data.id_sprzedazy} <br>
                            <strong>ID sklepu:</strong> {data.id_sklepu} <br>
                            <strong>Użytkownik:</strong> {data.id_uzytkownika} <br>
                            <strong>Kwota:</strong> {data.kwota_transakcji} zł <br>
                            <strong>Data:</strong> {formatDate(data.data_sprzedazy)} <br>
                        </div>
                    </li>
                {/each}
            </ul>
        </div>
    {/if}
</main>

<style>
    main {
        font-family: Arial, sans-serif;
        max-width: 600px;
        margin: 2em auto;
        padding: 1em;
        text-align: center;
    }

    form {
        display: flex;
        flex-direction: column;
        gap: 1em;
    }

    label {
        font-weight: bold;
    }

    select, input, button {
        padding: 0.5em;
        font-size: 1em;
        margin-top: 0.5em;
    }

    .error {
        color: red;
        margin-top: 1em;
    }

    .result {
        margin-top: 1em;
        text-align: left;
    }

    .result ul {
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
        text-align: left;
        width: 100%;
    }
</style>
