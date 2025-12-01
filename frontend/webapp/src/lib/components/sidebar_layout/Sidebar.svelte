<script lang="ts">
    let { children, userInfo = 'Admin User' } = $props();
</script>

<div id="layoutSidenav_nav">
    <nav class="sb-sidenav sb-sidenav-dark">
        <div class="sb-sidenav-menu">
            <div class="nav flex-column">
                {@render children?.()}
            </div>
        </div>

        <div class="sb-sidenav-footer">
            <div class="small">Logged in as:</div>
            {userInfo}
        </div>
    </nav>
</div>

<style>
    #layoutSidenav_nav {
        flex-basis: 225px;
        flex-shrink: 0;
        transition: transform 0.15s ease-in-out;
    }

    .sb-sidenav {
        display: flex;
        flex-direction: column;
        height: 100%;
        background-color: #212529;
    }

    .sb-sidenav-dark {
        background-color: #212529;
        color: rgba(255, 255, 255, 0.5);
    }

    .sb-sidenav-menu {
        flex-grow: 1;
        overflow-y: auto;
    }

    .sb-sidenav-footer {
        padding: 0.75rem 1rem;
        background-color: rgba(0, 0, 0, 0.1);
        font-size: 0.875rem;
    }

    /* Responsive - hide sidebar on mobile */
    @media (max-width: 991.98px) {
        #layoutSidenav_nav {
            transform: translateX(-225px);
        }
    }
</style>
