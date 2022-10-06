-- ################# NEW CRON JOB ##########################
CREATE OR REPLACE FUNCTION new_cron_job_notify()
	RETURNS trigger AS
$$
DECLARE
    w_id cron_workers.id%type;
BEGIN
    -- Choose insert cron job and assign worker to dispatch
    INSERT INTO cron_jobs(job_id, worker_id) VALUES (NEW.id,
        (SELECT id FROM cron_workers WHERE now() - updated_at < '2 second'::interval AND ready=true ORDER BY work_count ASC LIMIT 1)
    ) RETURNING worker_id INTO w_id;
    -- Notify worker
	PERFORM pg_notify(CONCAT('create_', w_id::text), NEW.id::text);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER new_cron_job
	AFTER INSERT
	ON jobs
	FOR EACH ROW
EXECUTE PROCEDURE new_cron_job_notify();

-- ################# DELETE CRON JOB ##########################
CREATE OR REPLACE FUNCTION delete_cron_job_notify()
	RETURNS trigger AS
$$
BEGIN
	PERFORM pg_notify(CONCAT('delete_', OLD.worker_id::text), OLD.job_id::text);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER delete_cron_job
	AFTER DELETE
	ON cron_jobs
	FOR EACH ROW
EXECUTE PROCEDURE delete_cron_job_notify();