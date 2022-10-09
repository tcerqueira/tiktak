export const fetchCronList = async () => {
    const response = await fetch("http://localhost:8080/cron");
    return response.json();
}

export interface PostCronPayload {
    webhook_url: string,
    webhook_method: string,
    body: string,
    cron_expression: string,
    timezone: string
}

export const postCron = async (payload: PostCronPayload) => {
    const response = fetch("http://localhost:8080/cron", {
        method: 'POST',
        mode: 'no-cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });
    return response;
}