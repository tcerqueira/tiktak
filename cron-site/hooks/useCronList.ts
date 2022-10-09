import { useState, useEffect, useCallback } from "react";
import { CronJob } from "../types/CronJob";
import { fetchCronList } from "../utils/ApiCalls";

export type CronListHookType = {
    cronList: CronJob[];
    isLoading: boolean;
    fetchList: () => void;
    error?: string;
}

export default function useCronList(): CronListHookType {
    const [cronList, setCronList] = useState<CronJob[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | undefined>();
    const [fetchFlag, setFetchFlag] = useState(false);

    const fetchList = useCallback(async () => {
        setFetchFlag(f => !f)
    }, []);

    useEffect(() => {
        (async () => {
            try {
                setIsLoading(true);
                const { data } = await fetchCronList();
                setIsLoading(false);
                setCronList(data);
            } catch (err) {
                console.log(err);
                setError('Failed to fecth data');
                return { cronList, isLoading, error, fetchList };
            }
        })();
    }, [fetchFlag]);

    return { cronList, isLoading, error, fetchList };
}