import { useState, useEffect } from "react";
import { CronJob } from "../types/CronJob";

export default function useCronList() {
    const [cronList, setCronList] = useState<CronJob[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>();
    const [error, setError] = useState<string>('');

    useEffect(() => {
        (async () => {
            try {
                setIsLoading(true);
                const response = await fetch("http://localhost:8080/cron");
                const { data } = await response.json();
                setIsLoading(false);
                setCronList(data);
            } catch (err) {
                console.log(err);
                setError('Failed to fecth data');
                return { cronList, isLoading, error };
            }
        })();
    }, []);

    return { cronList, isLoading, error };
}