CREATE KEYSPACE IF NOT EXISTS marketplace_analytics_stores
    WITH REPLICATION={'class':'SimpleStrategy', 'replication_factor':'3'};
