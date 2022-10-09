export interface EditCronPayload {
    body?: string,
    cron_expression?: string,
    timezone?: string
}

export interface PostCronPayload {
    webhook_url: string,
    webhook_method: string,
    body: string,
    cron_expression: string,
    timezone: string
}

export const fetchCronList = async () => {
    const response = await fetch("http://localhost:8080/cron");
    return response.json();
}

export const postCron = async (payload: PostCronPayload) => {
    const response = fetch("http://localhost:8080/cron", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload)
    });
    return response;
}

export const editCron = async (id: string, cron: EditCronPayload) => {
    const response = await fetch(`http://localhost:8080/cron/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(cron)
    });
    return response.json();
}

export const deleteCron = async (id: string) => {
    const response = fetch(`http://localhost:8080/cron/${id}`, {
        method: 'DELETE'
    });
    return response;
}
