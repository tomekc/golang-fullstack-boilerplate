<script lang="ts">
    import { TitleBar, HeroBar, Tile, Card, Table, Icon } from 'svebulma';

    const clients = [
        { name: 'Rebecca Bauch', company: 'Daugherty-Daniel', city: 'South Cory', progress: 79, created: 'Oct 25, 2020', avatar: 'rebecca-bauch' },
        { name: 'Felicita Yundt', company: 'Johns-Weissnat', city: 'East Ariel', progress: 67, created: 'Jan 8, 2020', avatar: 'felicita-yundt' },
        { name: 'Mr. Larry Satterfield V', company: 'Hyatt Ltd', city: 'Windlerburgh', progress: 16, created: 'Dec 18, 2020', avatar: 'mr-larry-satterfield-v' },
        { name: 'Mr. Broderick Kub', company: 'Kshlerin, Bauch and Ernser', city: 'New Kirstenport', progress: 71, created: 'Sep 13, 2020', avatar: 'mr-broderick-kub' },
        { name: 'Barry Weber', company: 'Schulist, Mosciski and Heidenreich', city: 'East Violettestad', progress: 80, created: 'Jul 24, 2020', avatar: 'barry-weber' },
    ];

    let currentPage = $state(1);
    const totalPages = 3;
</script>

<TitleBar breadcrumbs={['Admin', 'Dashboard']} />
<HeroBar title="Dashboard" />

<section class="section is-main-section">
    <div class="columns is-tiles-wrapper">
        <Tile label="Clients" value="512" icon="account-multiple" color="primary" />
        <Tile label="Sales" value="$7,770" icon="cart-outline" color="info" />
        <Tile label="Performance" value="256%" icon="finance" color="success" />
    </div>

    <Card title="Clients" icon="account-multiple" hasTable>
        {#snippet headerAction()}
            <Icon name="reload" />
        {/snippet}

        <Table>
            {#snippet head()}
                <tr>
                    <th></th>
                    <th>Name</th>
                    <th>Company</th>
                    <th>City</th>
                    <th>Progress</th>
                    <th>Created</th>
                    <th></th>
                </tr>
            {/snippet}

            {#snippet pagination()}
                <div class="level">
                    <div class="level-left">
                        <div class="level-item">
                            <div class="buttons has-addons">
                                {#each Array(totalPages) as _, i}
                                    <button
                                        type="button"
                                        class="button"
                                        class:is-active={currentPage === i + 1}
                                        onclick={() => currentPage = i + 1}
                                    >{i + 1}</button>
                                {/each}
                            </div>
                        </div>
                    </div>
                    <div class="level-right">
                        <div class="level-item">
                            <small>Page {currentPage} of {totalPages}</small>
                        </div>
                    </div>
                </div>
            {/snippet}

            {#each clients as client}
                <tr>
                    <td class="is-image-cell">
                        <div class="image">
                            <img src="https://avatars.dicebear.com/v2/initials/{client.avatar}.svg" alt={client.name} class="is-rounded">
                        </div>
                    </td>
                    <td data-label="Name">{client.name}</td>
                    <td data-label="Company">{client.company}</td>
                    <td data-label="City">{client.city}</td>
                    <td data-label="Progress" class="is-progress-cell">
                        <progress max="100" class="progress is-small is-primary" value={client.progress}>{client.progress}</progress>
                    </td>
                    <td data-label="Created">
                        <small class="has-text-grey is-abbr-like" title={client.created}>{client.created}</small>
                    </td>
                    <td class="is-actions-cell">
                        <div class="buttons is-right">
                            <button class="button is-small is-primary" type="button" aria-label="View">
                                <span class="icon"><i class="mdi mdi-eye"></i></span>
                            </button>
                            <button class="button is-small is-danger" type="button" aria-label="Delete">
                                <span class="icon"><i class="mdi mdi-trash-can"></i></span>
                            </button>
                        </div>
                    </td>
                </tr>
            {/each}
        </Table>
    </Card>
</section>

