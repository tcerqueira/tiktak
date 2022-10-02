export type CronJob= {
    id: number,
    webhook_url: string,
    webhook_method: string,
    body: string,
    expression: string,
    timezone: string,
    started_at?: string,
    last_trigger?: string,
};