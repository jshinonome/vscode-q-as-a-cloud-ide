.discovery.hosts: flip `host`port`label!"SJS"$\:();

upsert[`.discovery.hosts;(
  (`localhost;2000;`md.hk.rdb);
  (`localhost;2001;`md.hk.hdb);
  (`localhost;2002;`md.hk.tp)
 )];

.discovery.getHosts:{[user;password]
  .discovery.hosts
 };
