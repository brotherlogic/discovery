///
//  Generated code. Do not modify.
//  source: discovery.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const RegistryEntry$json = const {
  '1': 'RegistryEntry',
  '2': const [
    const {'1': 'ip', '3': 1, '4': 1, '5': 9, '10': 'ip'},
    const {'1': 'port', '3': 2, '4': 1, '5': 5, '10': 'port'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'external_port', '3': 4, '4': 1, '5': 8, '10': 'externalPort'},
    const {'1': 'identifier', '3': 5, '4': 1, '5': 9, '10': 'identifier'},
    const {'1': 'master', '3': 6, '4': 1, '5': 8, '10': 'master'},
    const {'1': 'weak_master', '3': 12, '4': 1, '5': 8, '10': 'weakMaster'},
    const {'1': 'register_time', '3': 7, '4': 1, '5': 3, '10': 'registerTime'},
    const {'1': 'time_to_clean', '3': 8, '4': 1, '5': 3, '10': 'timeToClean'},
    const {'1': 'last_seen_time', '3': 9, '4': 1, '5': 3, '10': 'lastSeenTime'},
    const {'1': 'ignores_master', '3': 10, '4': 1, '5': 8, '10': 'ignoresMaster'},
    const {'1': 'master_time', '3': 11, '4': 1, '5': 3, '10': 'masterTime'},
    const {'1': 'version', '3': 13, '4': 1, '5': 14, '6': '.discovery.RegistryEntry.Version', '10': 'version'},
  ],
  '4': const [RegistryEntry_Version$json],
};

const RegistryEntry_Version$json = const {
  '1': 'Version',
  '2': const [
    const {'1': 'V1', '2': 0},
    const {'1': 'V2', '2': 1},
  ],
};

const ServiceList$json = const {
  '1': 'ServiceList',
  '2': const [
    const {'1': 'services', '3': 1, '4': 3, '5': 11, '6': '.discovery.RegistryEntry', '10': 'services'},
  ],
};

const Empty$json = const {
  '1': 'Empty',
};

const StateResponse$json = const {
  '1': 'StateResponse',
  '2': const [
    const {'1': 'longest_call', '3': 1, '4': 1, '5': 3, '10': 'longestCall'},
    const {'1': 'most_frequent', '3': 2, '4': 1, '5': 9, '10': 'mostFrequent'},
    const {'1': 'frequency', '3': 3, '4': 1, '5': 5, '10': 'frequency'},
    const {'1': 'count', '3': 4, '4': 1, '5': 9, '10': 'count'},
  ],
};

const StateRequest$json = const {
  '1': 'StateRequest',
};

const RegisterRequest$json = const {
  '1': 'RegisterRequest',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
    const {'1': 'caller', '3': 2, '4': 1, '5': 9, '10': 'caller'},
    const {'1': 'fanout', '3': 4, '4': 1, '5': 8, '10': 'fanout'},
  ],
};

const RegisterResponse$json = const {
  '1': 'RegisterResponse',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
  ],
};

const DiscoverRequest$json = const {
  '1': 'DiscoverRequest',
  '2': const [
    const {'1': 'request', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'request'},
    const {'1': 'caller', '3': 2, '4': 1, '5': 9, '10': 'caller'},
  ],
};

const DiscoverResponse$json = const {
  '1': 'DiscoverResponse',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
  ],
};

const ListRequest$json = const {
  '1': 'ListRequest',
  '2': const [
    const {'1': 'caller', '3': 1, '4': 1, '5': 9, '10': 'caller'},
  ],
};

const ListResponse$json = const {
  '1': 'ListResponse',
  '2': const [
    const {'1': 'services', '3': 1, '4': 1, '5': 11, '6': '.discovery.ServiceList', '10': 'services'},
  ],
};

const GetRequest$json = const {
  '1': 'GetRequest',
  '2': const [
    const {'1': 'job', '3': 1, '4': 1, '5': 9, '10': 'job'},
    const {'1': 'server', '3': 2, '4': 1, '5': 9, '10': 'server'},
    const {'1': 'friend', '3': 3, '4': 1, '5': 9, '10': 'friend'},
  ],
};

const GetResponse$json = const {
  '1': 'GetResponse',
  '2': const [
    const {'1': 'services', '3': 1, '4': 3, '5': 11, '6': '.discovery.RegistryEntry', '10': 'services'},
  ],
};

const UnregisterRequest$json = const {
  '1': 'UnregisterRequest',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
    const {'1': 'address', '3': 3, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'caller', '3': 4, '4': 1, '5': 9, '10': 'caller'},
    const {'1': 'fanout', '3': 2, '4': 1, '5': 8, '10': 'fanout'},
  ],
};

const UnregisterResponse$json = const {
  '1': 'UnregisterResponse',
};

const LockRequest$json = const {
  '1': 'LockRequest',
  '2': const [
    const {'1': 'job', '3': 1, '4': 1, '5': 9, '10': 'job'},
    const {'1': 'lock_key', '3': 2, '4': 1, '5': 3, '10': 'lockKey'},
    const {'1': 'requestor', '3': 3, '4': 1, '5': 9, '10': 'requestor'},
  ],
};

const LockResponse$json = const {
  '1': 'LockResponse',
};

const MasterRequest$json = const {
  '1': 'MasterRequest',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
    const {'1': 'lock_key', '3': 2, '4': 1, '5': 3, '10': 'lockKey'},
    const {'1': 'master_elect', '3': 3, '4': 1, '5': 8, '10': 'masterElect'},
    const {'1': 'fanout', '3': 4, '4': 1, '5': 8, '10': 'fanout'},
  ],
};

const MasterResponse$json = const {
  '1': 'MasterResponse',
  '2': const [
    const {'1': 'service', '3': 1, '4': 1, '5': 11, '6': '.discovery.RegistryEntry', '10': 'service'},
  ],
};

