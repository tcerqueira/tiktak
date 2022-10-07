export type CronJob= {
    id: string,
    webhook_url: string,
    webhook_method: string,
    body: string,
    cron_expression: string,
    timezone: string,
    started_at?: string,
    last_trigger?: string,
};