export async function fetchUptimeData() {
    const response = await fetch('/api/monitor');
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}