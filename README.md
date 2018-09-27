# msvndr
Vendor version checker. It scans the entire dependency chain for all vendored
repos, and then repos that each vendored repo imports, continuing to the bottom
of the chain. It looks to see where repos have been vendored by multiple projects,
shows where there are discrepencies.

To install: `go get -u github.com/jhowardmsft/msvndr`

Here is an example (9/27/2018) run against docker/docker, containerd/cri and containerd/containerd. Note that the top-level repos are hard-coded in this code.

```
Analysing vendor.conf dependency chain:

   1: docker/docker
   2: Azure/go-ansiterm               d6e3b3328b783f23  docker/docker
   3: containerd/cri
   4: containerd/continuity           f44b615e492bdfb3  docker/docker
   5: Microsoft/hcsshim               v0.7.3            docker/docker
   6: Microsoft/go-winio              v0.4.11           docker/docker
   7: docker/libtrust                 9cbd2a1374f46905  docker/docker
   8: golang/gddo                     9b12a26f3fbd7397  docker/docker
   9: gorilla/context                 v1.1              docker/docker
  10: gorilla/mux                     v1.1              docker/docker
  11: Microsoft/opengcs               v0.3.9            docker/docker
  12: kr/pty                          5cf931ef8f        docker/docker
  13: mattn/go-shellwords             v1.0.3            docker/docker
  14: sirupsen/logrus                 v1.0.6            docker/docker
  15: tchap/go-patricia               v2.2.6            docker/docker
  16: vdemeester/shakers              24d7f1d6a71aa5d9  docker/docker
  17: google/go-cmp                   v0.2.0            docker/docker
  18: RackSec/srslog                  456df3a81436d29b  docker/docker
  19: imdario/mergo                   v0.3.6            docker/docker
  20: moby/buildkit                   39404586a50d1b9d  docker/docker
  21: tonistiigi/fsutil               b19464cd1b6a0077  docker/docker
  22: grpc-ecosystem/grpc-opentracin  8e809c8a86450a29  docker/docker
  23: opentracing/opentracing-go      1361b9cd60be79c4  docker/docker
  24: google/shlex                    6f45313302b9c568  docker/docker
  25: mitchellh/hashstructure         2bca23e0e452137f  docker/docker
  26: docker/libnetwork               36d3bed0e9f4b3c8  docker/docker
  27: docker/go-events                9461782956ad83b3  docker/docker
  28: armon/go-radix                  e39d623f12e8e41c  docker/docker
  29: armon/go-metrics                eb0af217e5e9747e  docker/docker
  30: hashicorp/go-msgpack            71c2886f5a673a35  docker/docker
  31: hashicorp/memberlist            3d8438da9589e7b6  docker/docker
  32: sean-/seed                      e2103e2c35297fb7  docker/docker
  33: hashicorp/go-sockaddr           6d291a969b86c4b6  docker/docker
  34: hashicorp/go-multierror         fcdddc395df1ddf4  docker/docker
  35: hashicorp/serf                  598c54895cc5a7b1  docker/docker
  36: docker/libkv                    458977154600b9f2  docker/docker
  37: vishvananda/netns               604eaf189ee867d8  docker/docker
  38: vishvananda/netlink             b2de5d10e38ecce8  docker/docker
  39: BurntSushi/toml                 a368813c5e648fee  docker/docker
  40: samuel/go-zookeeper             d0e0d8e11f318e00  docker/docker
  41: deckarep/golang-set             ef32fa3046d9f249  docker/docker
  42: coreos/etcd                     v3.3.9            docker/docker
  43: coreos/go-semver                v0.2.0            docker/docker
  44: ugorji/go                       v1.1.1            docker/docker
  45: hashicorp/consul                v0.5.2            docker/docker
  46: miekg/dns                       v1.0.7            docker/docker
  47: ishidawataru/sctp               07191f837fedd2f1  docker/docker
  48: docker/distribution             83389a148052d74a  docker/docker
  49: vbatts/tar-split                v0.11.0           docker/docker
  50: opencontainers/go-digest        v1.0.0-rc1        docker/docker
  51: mistifyio/go-zfs                22c9b32c84eb0d0c  docker/docker
  52: pborman/uuid                    v1.0              docker/docker
  53: opencontainers/runc             20aff4f0488c6d4b  docker/docker
  54: opencontainers/image-spec       v1.0.1            docker/docker
  55: seccomp/libseccomp-golang       32f571b70023028b  docker/docker
  56: coreos/go-systemd               v17               docker/docker
  57: godbus/dbus                     v4.0.0            docker/docker
  58: syndtr/gocapability             2c00daeb6c3b4511  docker/docker
  59: golang/protobuf                 v1.1.0            docker/docker
  60: Graylog2/go-gelf                4143646226541087  docker/docker
  61: fluent/fluent-logger-golang     v1.3.0            docker/docker
  62: philhofer/fwd                   98c11a7a6ec829d6  docker/docker
  63: tinylib/msgp                    3b556c64540842d4  docker/docker
  64: fsnotify/fsnotify               v1.4.7            docker/docker
  65: aws/aws-sdk-go                  v1.12.66          docker/docker
  66: go-ini/ini                      v1.25.4           docker/docker
  67: jmespath/go-jmespath            0b12d6b521d83fc7  docker/docker
  68: bsphere/le_go                   7a984a84b5492ae5  docker/docker
  69: googleapis/gax-go               v2.0.0            docker/docker
  70: containerd/containerd           d97a907f7f781c0a  docker/docker
  71: containerd/fifo                 3d5202aec260678c  docker/docker
  72: hashicorp/golang-lru            0fb14efe8c47ae85  docker/docker
  73: containerd/cgroups              5e610833b72089b3  docker/docker
  74: containerd/console              c12b1e7919c14469  docker/docker
  75: containerd/go-runc              5a6d9f37cfa36b15  docker/docker
  76: containerd/typeurl              a93fcdb778cd272c  docker/docker
  77: containerd/ttrpc                94dde388801693c5  docker/docker
  78: gogo/googleapis                 08a7655d27152912  docker/docker
  79: docker/swarmkit                 9f271c2963d18a7c  docker/docker
  80: gogo/protobuf                   v1.0.0            docker/docker
  81: cloudflare/cfssl                1.3.2             docker/docker
  82: fernet/fernet-go                1b2437bc582b3cfb  docker/docker
  83: google/certificate-transparenc  v1.0.20           docker/docker
  84: hashicorp/go-memdb              cb9a474f84cc5e41  docker/docker
  85: prometheus/procfs               7d6f385de8bea291  docker/docker
  86: coreos/pkg                      v3                docker/docker
  87: pivotal-golang/clock            3fd3c1944c59d974  docker/docker
  88: prometheus/client_golang        v0.8.0            docker/docker
  89: beorn7/perks                    3a771d992973f24a  docker/docker
  90: prometheus/client_model         6f38060186129309  docker/docker
  91: prometheus/common               7600349dcfe1abd1  docker/docker
  92: grpc-ecosystem/go-grpc-prometh  v1.2.0            docker/docker
  93: matttproud/golang_protobuf_ext  v1.0.0            docker/docker
  94: pkg/errors                      839d9e913e063e28  docker/docker
  95: spf13/pflag                     v1.0.1            docker/docker
  96: spf13/cobra                     v0.0.3            docker/docker
  97: docker/go-metrics               d466d4f6fd960e01  docker/docker
  98: opencontainers/selinux          b29023b86e4a69d1  docker/docker
  99: blang/semver                    v3.1.0            containerd/cri
 100: beorn7/perks                    4c0e84591b9aa9e6  containerd/cri
 101: containerd/cgroups              5e610833b72089b3  containerd/cri
 102: BurntSushi/toml                 a368813c5e648fee  containerd/cri
 103: containerd/continuity           7f53d412b9eb1cbf  containerd/cri
 104: containerd/console              c12b1e7919c14469  containerd/cri
 105: containerd/fifo                 3d5202aec260678c  containerd/cri
 106: containerd/go-runc              5a6d9f37cfa36b15  containerd/cri
 107: containerd/containerd           f88d3e5d6dfe9b7d  containerd/cri
 108: containerd/typeurl              a93fcdb778cd272c  containerd/cri
 109: containerd/go-cni               6d7b509a054a3cb1  containerd/cri
 110: containerd/ttrpc                2a805f7186350130  containerd/cri
 111: containernetworking/cni         v0.6.0            containerd/cri
 112: containernetworking/plugins     v0.7.0            containerd/cri
 113: coreos/go-systemd               v14               containerd/cri
 114: davecgh/go-spew                 v1.1.0            containerd/cri
 115: docker/docker                   86f080cff0914e96  containerd/cri
 116: docker/distribution             b38e5838b7b2f2ad  containerd/cri
 117: docker/go-metrics               4ea375f7759c8274  containerd/cri
 118: docker/go-events                9461782956ad83b3  containerd/cri
 119: docker/go-units                 v0.3.1            containerd/cri
 120: docker/spdystream               449fdfce4d962303  containerd/cri
 121: godbus/dbus                     v3                containerd/cri
 122: emicklei/go-restful             v2.2.1            containerd/cri
 123: ghodss/yaml                     v1.0.0            containerd/cri
 124: hashicorp/go-multierror         ed905158d8746222  containerd/cri
 125: gogo/googleapis                 08a7655d27152912  containerd/cri
 126: gogo/protobuf                   v1.0.0            containerd/cri
 127: golang/glog                     44145f04b68cf362  containerd/cri
 128: golang/protobuf                 v1.1.0            containerd/cri
 129: google/gofuzz                   44d81051d367757e  containerd/cri
 130: grpc-ecosystem/go-grpc-prometh  v1.1              containerd/cri
 131: hashicorp/errwrap               7554cd9344cec972  containerd/cri
 132: modern-go/reflect2              1.0.1             containerd/cri
 133: Microsoft/hcsshim               v0.7.4            containerd/cri
 134: modern-go/concurrent            1.0.3             containerd/cri
 135: json-iterator/go                1.1.5             containerd/cri
 136: matttproud/golang_protobuf_ext  v1.0.0            containerd/cri
 137: pkg/errors                      v0.8.0            containerd/cri
 138: opencontainers/go-digest        c9281466c8b2f606  containerd/cri
 139: opencontainers/image-spec       v1.0.1            containerd/cri
 140: opencontainers/runc             00dc70017d222b17  containerd/cri
 141: opencontainers/runtime-spec     eba862dc2470385a  containerd/cri
 142: opencontainers/runtime-tools    fb101d5d42ab9c04  containerd/cri
 143: opencontainers/selinux          b6fa367ed7f534f9  containerd/cri
 144: Microsoft/go-winio              v0.4.10           containerd/cri
 145: prometheus/client_golang        f4fb1b73fb099f39  containerd/cri
 146: prometheus/procfs               cb4147076ac75738  containerd/cri
 147: prometheus/client_model         99fa1f4be8e564e8  containerd/cri
 148: pmezard/go-difflib              v1.0.0            containerd/cri
 149: prometheus/common               89604d197083d478  containerd/cri
 150: syndtr/gocapability             db04d3cc01c8b549  containerd/cri
 151: urfave/cli                      7bc6a0acffa589f4  containerd/cri
 152: seccomp/libseccomp-golang       32f571b70023028b  containerd/cri
 153: containerd/containerd
 154: sirupsen/logrus                 v1.0.0            containerd/cri
 155: xeipuuv/gojsonschema            1d523034197ff1f2  containerd/cri
 156: xeipuuv/gojsonpointer           4e3ac2762d5f4793  containerd/cri
 157: tchap/go-patricia               v2.2.6            containerd/cri
 158: stretchr/testify                v1.1.4            containerd/cri
 159: xeipuuv/gojsonreference         bd5ef7bd5415a7ac  containerd/cri
 160: spf13/pflag                     4c012f6dcd954682  containerd/continuity
 161: inconshreveable/mousetrap       76626ae9c91c4f2a  containerd/continuity
 162: sirupsen/logrus                 89742aefa4b206dc  containerd/continuity
 163: golang/protobuf                 1e59b77b52bf8e4b  containerd/continuity
 164: spf13/cobra                     2da4a54c5ceefcee  containerd/continuity
 165: pkg/errors                      f15c970de5b76fac  containerd/continuity
 166: dustin/go-humanize              bb3d318650d48840  containerd/continuity
 167: opencontainers/go-digest        279bed98673dd5be  containerd/continuity
 168: Azure/azure-sdk-for-go          4650843026a7fdec  docker/distribution
 169: davecgh/go-spew                 v1.1.0            moby/buildkit
 170: Microsoft/hcsshim               v0.7.3            moby/buildkit
 171: seccomp/libseccomp-golang       32f571b70023028b  moby/buildkit
 172: Azure/go-autorest               eaa7994b2278094c  docker/distribution
 173: sirupsen/logrus                 3d4380f53a34dcdc  docker/distribution
 174: aws/aws-sdk-go                  f831d5a0822a1ad7  docker/distribution
 175: bugsnag/bugsnag-go              b1d153021fcd90ca  docker/distribution
 176: beorn7/perks                    4c0e84591b9aa9e6  docker/distribution
 177: denverdino/aliyungo             afedced274aa9a7f  docker/distribution
 178: marstr/guid                     8bd9a64bf37eb297  docker/distribution
 179: bugsnag/osext                   0dd3f918b21bec95  docker/distribution
 180: dgrijalva/jwt-go                a601269ab70c205d  docker/distribution
 181: docker/go-metrics               399ea8c73916000c  docker/distribution
 182: docker/libtrust                 fa567046d9b14f6a  docker/distribution
 183: garyburd/redigo                 535138d7bcd717d6  docker/distribution
 184: go-ini/ini                      2ba15ac2dc9cdf88  docker/distribution
 185: yvasiyarov/newrelic_platform_g  b21fdbd4370f3717  docker/distribution
 186: satori/go.uuid                  f58768cc1a7a7e77  docker/distribution
 187: matttproud/golang_protobuf_ext  c12348ce28de40ee  docker/distribution
 188: miekg/dns                       271c58e0c14f5521  docker/distribution
 189: bugsnag/panicwrap               e2c28503fcd06753  docker/distribution
 190: golang/protobuf                 8d92cf5fc15a4382  docker/distribution
 191: mitchellh/mapstructure          482a9fd5fa83e8c4  docker/distribution
 192: spf13/cobra                     312092086bed4968  docker/distribution
 193: inconshreveable/mousetrap       76626ae9c91c4f2a  docker/distribution
 194: gorilla/mux                     599cba5e7b6137d4  docker/distribution
 195: prometheus/procfs               cb4147076ac75738  docker/distribution
 196: ncw/swift                       a0320860b16212c2  docker/distribution
 197: spf13/pflag                     5644820622454e71  docker/distribution
 198: prometheus/client_golang        c332b6f63c0658a6  docker/distribution
 199: prometheus/client_model         99fa1f4be8e564e8  docker/distribution
 200: xenolf/lego                     a9d8cec0e6563575  docker/distribution
 201: yvasiyarov/go-metrics           57bccd1ccd43f94b  docker/distribution
 202: jmespath/go-jmespath            bd40a432e4c76585  docker/distribution
 203: gorilla/handlers                60c7bfde3e33c201  docker/distribution
 204: yvasiyarov/gorelic              a9bba5b9ab508a08  docker/distribution
 205: prometheus/common               89604d197083d478  docker/distribution
 206: urfave/cli                      d53eb991652b1d43  opencontainers/runc
 207: opencontainers/selinux          v1.0.0-rc1        opencontainers/runc
 208: seccomp/libseccomp-golang       84e90a91acea0f4e  opencontainers/runc
 209: sirupsen/logrus                 a3f95b5c42358657  opencontainers/runc
 210: syndtr/gocapability             db04d3cc01c8b549  opencontainers/runc
 211: vishvananda/netlink             1e2e08e8a2dcdaca  opencontainers/runc
 212: coreos/go-systemd               v14               opencontainers/runc
 213: coreos/pkg                      v3                opencontainers/runc
 214: godbus/dbus                     v3                opencontainers/runc
 215: golang/protobuf                 18c9bb3261723cd5  opencontainers/runc
 216: cyphar/filepath-securejoin      v0.2.1            opencontainers/runc
 217: docker/go-units                 v0.2.0            opencontainers/runc
 218: mrunalp/fileutils               ed869b029674c0e9  opencontainers/runc
 219: opencontainers/runtime-spec     v1.0.0            opencontainers/runc
 220: pkg/errors                      v0.8.0            moby/buildkit
 221: containerd/console              c12b1e7919c14469  moby/buildkit
 222: pmezard/go-difflib              v1.0.0            moby/buildkit
 223: containerd/containerd           d97a907f7f781c0a  moby/buildkit
 224: containerd/typeurl              a93fcdb778cd272c  moby/buildkit
 225: sirupsen/logrus                 v1.0.0            moby/buildkit
 226: opencontainers/go-digest        c9281466c8b2f606  moby/buildkit
 227: gogo/protobuf                   v1.0.0            moby/buildkit
 228: gogo/googleapis                 b23578765ee54ff6  moby/buildkit
 229: golang/protobuf                 v1.1.0            moby/buildkit
 230: containerd/continuity           f44b615e492bdfb3  moby/buildkit
 231: stretchr/testify                v1.1.4            moby/buildkit
 232: Microsoft/go-winio              v0.4.11           moby/buildkit
 233: opencontainers/image-spec       v1.0.1            moby/buildkit
 234: containerd/fifo                 3d5202aec260678c  moby/buildkit
 235: opencontainers/runc             20aff4f0488c6d4b  moby/buildkit
 236: containerd/go-runc              5a6d9f37cfa36b15  moby/buildkit
 237: opentracing-contrib/go-stdlib   b1a47cfbdd7543e7  moby/buildkit
 238: docker/distribution             30578ca32960a4d3  moby/buildkit
 239: tonistiigi/units                6950e57a87eaf136  moby/buildkit
 240: docker/docker-credential-helpe  d68f9aeca33f5fd3  moby/buildkit
 241: docker/libnetwork               36d3bed0e9f4b3c8  moby/buildkit
 242: BurntSushi/toml                 3012a1dbe2e4bd13  moby/buildkit
 243: ishidawataru/sctp               07191f837fedd2f1  moby/buildkit
 244: grpc-ecosystem/grpc-opentracin  8e809c8a86450a29  moby/buildkit
 245: opentracing/opentracing-go      1361b9cd60be79c4  moby/buildkit
 246: uber/jaeger-client-go           e02c85f9069ea625  moby/buildkit
 247: apache/thrift                   b2a4d4ae21c789b6  moby/buildkit
 248: uber/jaeger-lib                 c48167d9cae58873  moby/buildkit
 249: codahale/hdrhistogram           f8ad88b59a584afe  moby/buildkit
 250: docker/go-events                9461782956ad83b3  moby/buildkit
 251: bshuster-repo/logrus-logstash-  d2c0ecc1836d9181  docker/distribution
 252: docker/docker                   71cd53e4a197b303  moby/buildkit
 253: urfave/cli                      7bc6a0acffa589f4  moby/buildkit
 254: morikuni/aec                    39771216ff4c63d1  moby/buildkit
 255: docker/go-units                 v0.3.1            moby/buildkit
 256: google/shlex                    6f45313302b9c568  moby/buildkit
 257: tonistiigi/fsutil               7e391b0e788f9b92  moby/buildkit
 258: pkg/profile                     5b67d428864e9271  moby/buildkit
 259: hashicorp/golang-lru            a0d98a5f28801957  moby/buildkit
 260: mitchellh/hashstructure         2bca23e0e452137f  moby/buildkit
 261: syndtr/gocapability             db04d3cc01c8b549  moby/buildkit
 262: docker/go-connections           3ede32e2033de750  moby/buildkit
 263: google/go-cmp                   v0.2.0            moby/buildkit
 264: deckarep/golang-set             ef32fa3046d9f249  docker/libnetwork
 265: armon/go-radix                  e39d623f12e8e41c  docker/libnetwork
 266: codegangsta/cli                 a65b733b303f0055  docker/libnetwork
 267: containerd/continuity           d3c23511c1bf5851  docker/libnetwork
 268: armon/go-metrics                eb0af217e5e9747e  docker/libnetwork
 269: Azure/go-ansiterm               d6e3b3328b783f23  docker/libnetwork
 270: coreos/etcd                     v3.2.1            docker/libnetwork
 271: coreos/go-semver                v0.2.0            docker/libnetwork
 272: BurntSushi/toml                 a368813c5e648fee  docker/libnetwork
 273: Microsoft/go-winio              v0.4.11           docker/libnetwork
 274: Microsoft/hcsshim               v0.7.3            docker/libnetwork
 275: containerd/fifo                 3d5202aec260678c  containerd/containerd
 276: containerd/go-runc              5a6d9f37cfa36b15  containerd/containerd
 277: containerd/console              c12b1e7919c14469  containerd/containerd
 278: containerd/cgroups              5e610833b72089b3  containerd/containerd
 279: containerd/typeurl              a93fcdb778cd272c  containerd/containerd
 280: golang/protobuf                 v1.1.0            containerd/containerd
 281: containerd/btrfs                2e1aa0ddf94f91fa  containerd/containerd
 282: prometheus/client_model         99fa1f4be8e564e8  containerd/containerd
 283: prometheus/client_golang        f4fb1b73fb099f39  containerd/containerd
 284: prometheus/common               89604d197083d478  containerd/containerd
 285: containerd/continuity           7f53d412b9eb1cbf  containerd/containerd
 286: coreos/go-systemd               48702e0da86bd25e  containerd/containerd
 287: prometheus/procfs               cb4147076ac75738  containerd/containerd
 288: docker/go-metrics               4ea375f7759c8274  containerd/containerd
 289: docker/go-events                9461782956ad83b3  containerd/containerd
 290: beorn7/perks                    4c0e84591b9aa9e6  containerd/containerd
 291: docker/go-units                 v0.3.1            containerd/containerd
 292: godbus/dbus                     c7fdd8b5cd55e87b  containerd/containerd
 293: matttproud/golang_protobuf_ext  v1.0.0            containerd/containerd
 294: gogo/protobuf                   v1.0.0            containerd/containerd
 295: gogo/googleapis                 08a7655d27152912  containerd/containerd
 296: pkg/errors                      v0.8.0            containerd/go-cni
 297: stretchr/testify                b89eecf5ca5db6d3  containerd/go-cni
 298: davecgh/go-spew                 8991bc29aa16c548  containerd/go-cni
 299: pmezard/go-difflib              792786c7400a1362  containerd/go-cni
 300: stretchr/objx                   8a3f7159479fbc75  containerd/go-cni
 301: containernetworking/cni         142cde0c766cd605  containerd/go-cni
 302: pkg/errors                      v0.8.0            opencontainers/runc
 303: containerd/console              2748ece16665b45a  opencontainers/runc
 304: vishvananda/netns               604eaf189ee867d8  docker/libnetwork
 305: hashicorp/memberlist            3d8438da9589e7b6  docker/libnetwork
 306: docker/docker                   162ba6016def6726  docker/libnetwork
 307: gorilla/context                 v1.1              docker/libnetwork
 308: gorilla/mux                     v1.1              docker/libnetwork
 309: hashicorp/consul                v0.5.2            docker/libnetwork
 310: hashicorp/go-msgpack            71c2886f5a673a35  docker/libnetwork
 311: hashicorp/go-multierror         fcdddc395df1ddf4  docker/libnetwork
 312: docker/go-units                 9e638d38cf6977a3  docker/libnetwork
 313: docker/go-connections           7beb39f0b969b075  docker/libnetwork
 314: docker/go-events                9461782956ad83b3  docker/libnetwork
 315: opencontainers/image-spec       v1.0.1            docker/libnetwork
 316: sean-/seed                      e2103e2c35297fb7  docker/libnetwork
 317: hashicorp/go-sockaddr           6d291a969b86c4b6  docker/libnetwork
 318: hashicorp/serf                  598c54895cc5a7b1  docker/libnetwork
 319: mattn/go-shellwords             v1.0.3            docker/libnetwork
 320: miekg/dns                       v1.0.7            docker/libnetwork
 321: opencontainers/go-digest        v1.0.0-rc1        docker/libnetwork
 322: samuel/go-zookeeper             d0e0d8e11f318e00  docker/libnetwork
 323: opencontainers/runc             69663f0bd4b60df0  docker/libnetwork
 324: opencontainers/runtime-spec     v1.0.1            docker/libnetwork
 325: docker/libkv                    458977154600b9f2  docker/libnetwork
 326: ugorji/go                       f1f1a805ed361a0e  docker/libnetwork
 327: sirupsen/logrus                 v1.0.3            docker/libnetwork
 328: vishvananda/netlink             b2de5d10e38ecce8  docker/libnetwork
 329: godbus/dbus                     v4.0.0            docker/libnetwork
 330: gogo/protobuf                   v1.0.0            docker/libnetwork
 331: urfave/cli                      7bc6a0acffa589f4  containerd/containerd
 332: opencontainers/runc             00dc70017d222b17  containerd/containerd
 333: sirupsen/logrus                 v1.0.0            containerd/containerd
 334: matttproud/golang_protobuf_ext  v1.0.0            docker/swarmkit
 335: gogo/protobuf                   v1.0.0            docker/swarmkit
 336: golang/protobuf                 v1.1.0            docker/swarmkit
 337: ishidawataru/sctp               07191f837fedd2f1  docker/swarmkit
 338: docker/distribution             83389a148052d74a  docker/swarmkit
 339: grpc-ecosystem/go-grpc-prometh  6b7015e65d366bf3  docker/swarmkit
 340: docker/go-metrics               d466d4f6fd960e01  docker/swarmkit
 341: coreos/etcd                     v3.2.1            docker/swarmkit
 342: coreos/go-systemd               v17               docker/swarmkit
 343: coreos/pkg                      v3                docker/swarmkit
 344: prometheus/client_golang        52437c81da6b127a  docker/swarmkit
 345: prometheus/client_model         fa8ad6fec33561be  docker/swarmkit
 346: prometheus/common               ebdfc6da46522d58  docker/swarmkit
 347: prometheus/procfs               abf152e5f3e97f2f  docker/swarmkit
 348: docker/libkv                    1d8431073ae03cda  docker/swarmkit
 349: docker/docker                   b9bb3bae5161f931  docker/swarmkit
 350: docker/go-connections           7beb39f0b969b075  docker/swarmkit
 351: docker/go-events                9461782956ad83b3  docker/swarmkit
 352: docker/go-units                 9e638d38cf6977a3  docker/swarmkit
 353: opencontainers/runc             ad0f5255060d3687  docker/swarmkit
 354: docker/libnetwork               a79d368793169724  docker/swarmkit
 355: opencontainers/go-digest        v1.0.0-rc1        docker/swarmkit
 356: opencontainers/image-spec       v1.0.1            docker/swarmkit
 357: opencontainers/go-digest        c9281466c8b2f606  containerd/containerd
 358: google/certificate-transparenc  v1.0.20           docker/swarmkit
 359: pkg/errors                      v0.8.0            containerd/containerd
 360: beorn7/perks                    4c0e84591b9aa9e6  docker/swarmkit
 361: Microsoft/go-winio              v0.4.8            docker/swarmkit
 362: sirupsen/logrus                 v1.0.3            docker/swarmkit
 363: dustin/go-humanize              8929fe90cee4b2cb  docker/swarmkit
 364: cloudflare/cfssl                1.3.2             docker/swarmkit
 365: fernet/fernet-go                1b2437bc582b3cfb  docker/swarmkit
 366: opencontainers/image-spec       v1.0.1            containerd/containerd
 367: stretchr/testify                v1.1.4            docker/swarmkit
 368: pivotal-golang/clock            3fd3c1944c59d974  docker/swarmkit
 369: hashicorp/go-memdb              cb9a474f84cc5e41  docker/swarmkit
 370: hashicorp/golang-lru            a0d98a5f28801957  docker/swarmkit
 371: inconshreveable/mousetrap       76626ae9c91c4f2a  docker/swarmkit
 372: phayes/permbits                 f7e3ac5e859d0b91  docker/swarmkit
 373: rcrowley/go-metrics             51425a2415d21afa  docker/swarmkit
 374: pkg/errors                      645ef00459ed84a1  docker/swarmkit
 375: pmezard/go-difflib              792786c7400a1362  docker/swarmkit
 376: spf13/cobra                     8e91712f174ced10  docker/swarmkit
 377: spf13/pflag                     7f60f83a2c81bc3c  docker/swarmkit
 378: ishidawataru/sctp               07191f837fedd2f1  docker/libnetwork
 379: Microsoft/hcsshim               v0.7.4            containerd/containerd
 380: pkg/errors                      839d9e913e063e28  docker/libnetwork
 381: grpc-ecosystem/go-grpc-prometh  6b7015e65d366bf3  containerd/containerd
 382: BurntSushi/toml                 a368813c5e648fee  containerd/containerd
 383: Microsoft/go-winio              v0.4.10           containerd/containerd
 384: pkg/errors                      v0.8.0            cyphar/filepath-securejoin
 385: google/go-cmp                   v0.2.0            docker/libnetwork
 386: opencontainers/image-spec       ab7389ef9f50030c  docker/distribution
 387: containerd/aufs                 ffa39970e26ad01d  containerd/containerd
 388: golang/glog                     44145f04b68cf362  containerd/containerd
 389: containerd/ttrpc                2a805f7186350130  containerd/containerd
 390: opencontainers/go-digest        a6d0ee40d4207ea0  docker/distribution
 391: syndtr/gocapability             db04d3cc01c8b549  containerd/containerd
 392: google/go-cmp                   v0.1.0            containerd/containerd
 393: containerd/go-cni               6d7b509a054a3cb1  containerd/containerd
 394: blang/semver                    v3.1.0            containerd/containerd
 395: containernetworking/cni         v0.6.0            containerd/containerd
 396: containernetworking/plugins     v0.7.0            containerd/containerd
 397: davecgh/go-spew                 v1.1.0            containerd/containerd
 398: docker/distribution             b38e5838b7b2f2ad  containerd/containerd
 399: docker/docker                   86f080cff0914e96  containerd/containerd
 400: docker/spdystream               449fdfce4d962303  containerd/containerd
 401: emicklei/go-restful             v2.2.1            containerd/containerd
 402: ghodss/yaml                     v1.0.0            containerd/containerd
 403: opencontainers/selinux          b6fa367ed7f534f9  containerd/containerd
 404: google/gofuzz                   44d81051d367757e  containerd/containerd
 405: hashicorp/errwrap               7554cd9344cec972  containerd/containerd
 406: hashicorp/go-multierror         ed905158d8746222  containerd/containerd
 407: json-iterator/go                1.1.5             containerd/containerd
 408: modern-go/reflect2              1.0.1             containerd/containerd
 409: modern-go/concurrent            1.0.3             containerd/containerd
 410: opencontainers/runtime-tools    v0.6.0            containerd/containerd
 411: containerd/zfs                  9a0b8b8b5982014b  containerd/containerd
 412: xeipuuv/gojsonschema            1d523034197ff1f2  containerd/containerd
 413: tchap/go-patricia               v2.2.6            containerd/containerd
 414: seccomp/libseccomp-golang       32f571b70023028b  containerd/containerd
 415: xeipuuv/gojsonpointer           4e3ac2762d5f4793  containerd/containerd
 416: mistifyio/go-zfs                166add352731e515  containerd/containerd
 417: pborman/uuid                    c65b2f87fee37d1c  containerd/containerd
 418: xeipuuv/gojsonreference         bd5ef7bd5415a7ac  containerd/containerd


The following repos were skipped (either aliased or not under github.com):

        - github.com/go-check/check
        - golang.org/x/net
        - golang.org/x/sys
        - github.com/docker/go-units
        - github.com/docker/go-connections
        - golang.org/x/text
        - gotest.tools
        - golang.org/x/sync
        - github.com/opentracing-contrib/go-stdlib
        - go.etcd.io/bbolt
        - google.golang.org/grpc
        - github.com/opencontainers/runtime-spec
        - golang.org/x/oauth2
        - google.golang.org/api
        - go.opencensus.io
        - cloud.google.com/go
        - google.golang.org/genproto
        - golang.org/x/crypto
        - golang.org/x/time
        - github.com/hashicorp/go-immutable-radix
        - github.com/inconshreveable/mousetrap
        - github.com/Nvveen/Gotty
        - go.etcd.io/bbolt
        - golang.org/x/crypto
        - golang.org/x/net
        - golang.org/x/oauth2
        - golang.org/x/sync
        - golang.org/x/sys
        - golang.org/x/text
        - golang.org/x/time
        - google.golang.org/genproto
        - google.golang.org/grpc
        - gopkg.in/inf.v0
        - gopkg.in/yaml.v2
        - k8s.io/api
        - k8s.io/apimachinery
        - k8s.io/apiserver
        - k8s.io/client-go
        - k8s.io/kubernetes
        - k8s.io/utils
        - bazil.org/fuse
        - golang.org/x/crypto
        - golang.org/x/net
        - golang.org/x/sync
        - golang.org/x/sys
        - go.etcd.io/bbolt
        - golang.org/x/sys
        - golang.org/x/sync
        - google.golang.org/grpc
        - golang.org/x/net
        - github.com/opencontainers/runtime-spec
        - google.golang.org/genproto
        - golang.org/x/text
        - golang.org/x/crypto
        - golang.org/x/time
        - github.com/hashicorp/go-immutable-radix
        - github.com/docker/cli
        - gotest.tools
        - golang.org/x/crypto
        - golang.org/x/sys
        - go.etcd.io/bbolt
        - github.com/opencontainers/runtime-spec
        - google.golang.org/grpc
        - golang.org/x/net
        - golang.org/x/crypto
        - golang.org/x/net
        - google.golang.org/genproto
        - golang.org/x/oauth2
        - golang.org/x/net
        - google.golang.org/grpc
        - github.com/davecgh/go-spew
        - golang.org/x/time
        - golang.org/x/sys
        - golang.org/x/sys
        - github.com/hashicorp/go-immutable-radix
        - google.golang.org/api
        - golang.org/x/sync
        - golang.org/x/sync
        - go.etcd.io/bbolt
        - google.golang.org/appengine
        - gotest.tools
        - google.golang.org/genproto
        - golang.org/x/crypto
        - google.golang.org/cloud
        - golang.org/x/text
        - golang.org/x/net
        - golang.org/x/sys
        - golang.org/x/text
        - golang.org/x/time
        - google.golang.org/grpc
        - gopkg.in/check.v1
        - gopkg.in/square/go-jose.v1
        - gopkg.in/yaml.v2
        - rsc.io/letsencrypt
        - gotest.tools
        - go.etcd.io/bbolt
        - github.com/containerd/cri
        - golang.org/x/crypto
        - golang.org/x/oauth2
        - golang.org/x/time
        - gopkg.in/inf.v0
        - gopkg.in/yaml.v2
        - k8s.io/api
        - k8s.io/apimachinery
        - k8s.io/apiserver
        - k8s.io/client-go
        - k8s.io/kubernetes
        - k8s.io/utils


Analysing the results:


WARN: hashicorp/go-multierror has 2 versions imported
        fcdddc395df1ddf4247c69bd436e84cfa0733f7e by:
                docker/docker
                docker/libnetwork
        ed905158d87462226a13fe39ddf685ea65f1c11f by:
                containerd/cri
                containerd/containerd


WARN: mistifyio/go-zfs has 2 versions imported
        22c9b32c84eb0d0c6f4043b6e90fc94073de92fa by:
                docker/docker
        166add352731e515512690329794ee593f1aaff2 by:
                containerd/containerd


WARN: prometheus/client_model has 3 versions imported
        99fa1f4be8e564e8a6b613da7fa6f46c9edafc6c by:
                containerd/cri
                docker/distribution
                containerd/containerd
        fa8ad6fec33561be4280a8f0514318c79d7f6cb6 by:
                docker/swarmkit
        6f3806018612930941127f2a7c6c453ba2c527d2 by:
                docker/docker


WARN: golang/protobuf has 4 versions imported
        v1.1.0 by:
                docker/docker
                containerd/cri
                moby/buildkit
                containerd/containerd
                docker/swarmkit
        1e59b77b52bf8e4b449a57e6f79f21226d571845 by:
                containerd/continuity
        8d92cf5fc15a4382f8964b08e1f42a75c0591aa3 by:
                docker/distribution
        18c9bb3261723cd5401db4d0c9fbc5c3b6c70fe8 by:
                opencontainers/runc


WARN: BurntSushi/toml has 2 versions imported
        a368813c5e648fee92e5f6c30e3944ff9d5e8895 by:
                docker/docker
                containerd/cri
                docker/libnetwork
                containerd/containerd
        3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005 by:
                moby/buildkit


WARN: sirupsen/logrus has 6 versions imported
        89742aefa4b206dcf400792f3bd35b542998eb3b by:
                containerd/continuity
        3d4380f53a34dcdc95f0c1db702615992b38d9a4 by:
                docker/distribution
        a3f95b5c423586578a4e099b11a46c2479628cac by:
                opencontainers/runc
        v1.0.3 by:
                docker/libnetwork
                docker/swarmkit
        v1.0.6 by:
                docker/docker
        v1.0.0 by:
                containerd/cri
                moby/buildkit
                containerd/containerd


WARN: ugorji/go has 2 versions imported
        v1.1.1 by:
                docker/docker
        f1f1a805ed361a0e078bb537e4ea78cd37dcf065 by:
                docker/libnetwork


WARN: hashicorp/golang-lru has 2 versions imported
        0fb14efe8c47ae851c0034ed7a448854d3d34cf3 by:
                docker/docker
        a0d98a5f288019575c6d1f4bb1573fef2d1fcdc4 by:
                moby/buildkit
                docker/swarmkit


WARN: dustin/go-humanize has 2 versions imported
        bb3d318650d48840a39aa21a027c6630e198e626 by:
                containerd/continuity
        8929fe90cee4b2cb9deb468b51fb34eba64d1bf0 by:
                docker/swarmkit


WARN: seccomp/libseccomp-golang has 2 versions imported
        32f571b70023028bd57d9288c20efbcb237f3ce0 by:
                docker/docker
                containerd/cri
                moby/buildkit
                containerd/containerd
        84e90a91acea0f4e51e62bc1a75de18b1fc0790f by:
                opencontainers/runc


WARN: beorn7/perks has 2 versions imported
        3a771d992973f24aa725d07868b467d1ddfceaf by:
                docker/docker
        4c0e84591b9aa9e6dcfdf3e020114cd81f89d5f9 by:
                containerd/cri
                docker/distribution
                containerd/containerd
                docker/swarmkit


WARN: opencontainers/runtime-tools has 2 versions imported
        fb101d5d42ab9c040f7d0a004e78336e5d5cb197 by:
                containerd/cri
        v0.6.0 by:
                containerd/containerd


WARN: docker/docker has 5 versions imported
         by:

        86f080cff0914e9694068ed78d503701667c4c00 by:
                containerd/cri
                containerd/containerd
        71cd53e4a197b303c6ba086bd584ffd67a884281 by:
                moby/buildkit
        162ba6016def672690ee4a1f3978368853a1e149 by:
                docker/libnetwork
        b9bb3bae5161f931c1dede43c67948c599197f50 by:
                docker/swarmkit


WARN: coreos/go-systemd has 3 versions imported
        48702e0da86bd25e76cfef347e2adeb434a0d0a6 by:
                containerd/containerd
        v17 by:
                docker/docker
                docker/swarmkit
        v14 by:
                containerd/cri
                opencontainers/runc


WARN: aws/aws-sdk-go has 2 versions imported
        v1.12.66 by:
                docker/docker
        f831d5a0822a1ad72420ab18c6269bca1ddaf490 by:
                docker/distribution


WARN: stretchr/testify has 2 versions imported
        v1.1.4 by:
                containerd/cri
                moby/buildkit
                docker/swarmkit
        b89eecf5ca5db6d3ba60b237ffe3df7bafb7662f by:
                containerd/go-cni


WARN: godbus/dbus has 3 versions imported
        v4.0.0 by:
                docker/docker
                docker/libnetwork
        v3 by:
                containerd/cri
                opencontainers/runc
        c7fdd8b5cd55e87b4e1f4e372cdb1db61dd6c66f by:
                containerd/containerd


WARN: prometheus/procfs has 3 versions imported
        cb4147076ac75738c9a7d279075a253c0cc5acbd by:
                containerd/cri
                docker/distribution
                containerd/containerd
        abf152e5f3e97f2fafac028d2cc06c1feb87ffa5 by:
                docker/swarmkit
        7d6f385de8bea29190f15ba9931442a0eaef9af7 by:
                docker/docker


WARN: pkg/errors has 4 versions imported
        645ef00459ed84a119197bfb8d8205042c6df63d by:
                docker/swarmkit
        839d9e913e063e28dfd0e6c7b7512793e0a48be9 by:
                docker/docker
                docker/libnetwork
        v0.8.0 by:
                containerd/cri
                moby/buildkit
                containerd/go-cni
                opencontainers/runc
                containerd/containerd
                cyphar/filepath-securejoin
        f15c970de5b76fac0b59abb32d62c17cc7bed265 by:
                containerd/continuity


WARN: docker/go-units has 3 versions imported
        v0.3.1 by:
                containerd/cri
                moby/buildkit
                containerd/containerd
        v0.2.0 by:
                opencontainers/runc
        9e638d38cf6977a37a8ea0078f3ee75a7cdb2dd1 by:
                docker/libnetwork
                docker/swarmkit


WARN: opencontainers/runtime-spec has 3 versions imported
        eba862dc2470385a233c7507392675cbeadf7353 by:
                containerd/cri
        v1.0.0 by:
                opencontainers/runc
        v1.0.1 by:
                docker/libnetwork


WARN: go-ini/ini has 2 versions imported
        v1.25.4 by:
                docker/docker
        2ba15ac2dc9cdf88c110ec2dc0ced7fa45f5678c by:
                docker/distribution


WARN: jmespath/go-jmespath has 2 versions imported
        0b12d6b521d83fc7f755e7cfc1b1fbdd35a01a74 by:
                docker/docker
        bd40a432e4c76585ef6b72d3fd96fb9b6dc7b68d by:
                docker/distribution


WARN: google/go-cmp has 2 versions imported
        v0.1.0 by:
                containerd/containerd
        v0.2.0 by:
                docker/docker
                moby/buildkit
                docker/libnetwork


WARN: docker/libnetwork has 2 versions imported
        36d3bed0e9f4b3c8c66df9bd45278bb90b33e911 by:
                docker/docker
                moby/buildkit
        a79d3687931697244b8e03485bf7b2042f8ec6b6 by:
                docker/swarmkit


WARN: davecgh/go-spew has 2 versions imported
        v1.1.0 by:
                containerd/cri
                moby/buildkit
                containerd/containerd
        8991bc29aa16c548c550c7ff78260e27b9ab7c73 by:
                containerd/go-cni


WARN: vishvananda/netlink has 2 versions imported
        b2de5d10e38ecce8607e6b438b6d174f389a004e by:
                docker/docker
                docker/libnetwork
        1e2e08e8a2dcdacaae3f14ac44c5cfa31361f270 by:
                opencontainers/runc


WARN: tonistiigi/fsutil has 2 versions imported
        b19464cd1b6a00773b4f2eb7acf9c30426f9df42 by:
                docker/docker
        7e391b0e788f9b925f22bd3cf88e0210d1643673 by:
                moby/buildkit


WARN: containerd/continuity has 3 versions imported
        f44b615e492bdfb371aae2f76ec694d9da1db537 by:
                docker/docker
                moby/buildkit
        7f53d412b9eb1cbf744c2063185d703a0ee34700 by:
                containerd/cri
                containerd/containerd
        d3c23511c1bf5851696cba83143d9cbcd666869b by:
                docker/libnetwork


WARN: containerd/console has 2 versions imported
        c12b1e7919c14469339a5d38f2f8ed9b64a9de23 by:
                docker/docker
                containerd/cri
                moby/buildkit
                containerd/containerd
        2748ece16665b45a47f884001d5831ec79703880 by:
                opencontainers/runc


WARN: docker/libtrust has 2 versions imported
        fa567046d9b14f6aa788882a950d69651d230b21 by:
                docker/distribution
        9cbd2a1374f46905c68a4eb3694a130610adc62a by:
                docker/docker


WARN: opencontainers/runc has 4 versions imported
        69663f0bd4b60df09991c08812a60108003fa340 by:
                docker/libnetwork
        ad0f5255060d36872be04de22f8731f38ef2d7b1 by:
                docker/swarmkit
        20aff4f0488c6d4b8df4d85b4f63f1f704c11abd by:
                docker/docker
                moby/buildkit
        00dc70017d222b178a002ed30e9321b12647af2d by:
                containerd/cri
                containerd/containerd


WARN: pmezard/go-difflib has 2 versions imported
        v1.0.0 by:
                containerd/cri
                moby/buildkit
        792786c7400a136282c1664665ae0a8db921c6c2 by:
                containerd/go-cni
                docker/swarmkit


WARN: Microsoft/hcsshim has 2 versions imported
        v0.7.3 by:
                docker/docker
                moby/buildkit
                docker/libnetwork
        v0.7.4 by:
                containerd/cri
                containerd/containerd


WARN: containerd/ttrpc has 2 versions imported
        94dde388801693c54f88a6596f713b51a8b30b2d by:
                docker/docker
        2a805f71863501300ae1976d29f0454ae003e85a by:
                containerd/cri
                containerd/containerd


WARN: prometheus/client_golang has 4 versions imported
        52437c81da6b127a9925d17eb3a382a2e5fd395e by:
                docker/swarmkit
        v0.8.0 by:
                docker/docker
        f4fb1b73fb099f396a7f0036bf86aa8def4ed823 by:
                containerd/cri
                containerd/containerd
        c332b6f63c0658a65eca15c0e5247ded801cf564 by:
                docker/distribution


WARN: opencontainers/selinux has 3 versions imported
        b29023b86e4a69d1b46b7e7b4e2b6fda03f0b9cd by:
                docker/docker
        b6fa367ed7f534f9ba25391cc2d467085dbb445a by:
                containerd/cri
                containerd/containerd
        v1.0.0-rc1 by:
                opencontainers/runc


WARN: spf13/cobra has 4 versions imported
        312092086bed4968099259622145a0c9ae280064 by:
                docker/distribution
        8e91712f174ced10270cf66615e0a9127e7c4de5 by:
                docker/swarmkit
        v0.0.3 by:
                docker/docker
        2da4a54c5ceefcee7ca5dd0eea1e18a3b6366489 by:
                containerd/continuity


WARN: docker/go-connections has 2 versions imported
        3ede32e2033de7505e6500d6c868c2b9ed9f169d by:
                moby/buildkit
        7beb39f0b969b075d1325fecb092faf27fd357b6 by:
                docker/libnetwork
                docker/swarmkit


WARN: containernetworking/cni has 2 versions imported
        v0.6.0 by:
                containerd/cri
                containerd/containerd
        142cde0c766cd6055cc7fdfdcb44579c0c9c35bf by:
                containerd/go-cni


WARN: miekg/dns has 2 versions imported
        v1.0.7 by:
                docker/docker
                docker/libnetwork
        271c58e0c14f552178ea321a545ff9af38930f39 by:
                docker/distribution


WARN: opencontainers/go-digest has 4 versions imported
        c9281466c8b2f606084ac71339773efd177436e7 by:
                containerd/cri
                moby/buildkit
                containerd/containerd
        279bed98673dd5bef374d3b6e4b09e2af76183bf by:
                containerd/continuity
        a6d0ee40d4207ea02364bd3b9e8e77b9159ba1eb by:
                docker/distribution
        v1.0.0-rc1 by:
                docker/docker
                docker/libnetwork
                docker/swarmkit


WARN: Microsoft/go-winio has 3 versions imported
        v0.4.11 by:
                docker/docker
                moby/buildkit
                docker/libnetwork
        v0.4.10 by:
                containerd/cri
                containerd/containerd
        v0.4.8 by:
                docker/swarmkit


WARN: gorilla/mux has 2 versions imported
        v1.1 by:
                docker/docker
                docker/libnetwork
        599cba5e7b6137d46ddf58fb1765f5d928e69604 by:
                docker/distribution


WARN: containerd/containerd has 3 versions imported
        d97a907f7f781c0ab8340877d8e6b53cc7f1c2f6 by:
                docker/docker
                moby/buildkit
        f88d3e5d6dfe9b7d7941ac5241649ad8240b9282 by:
                containerd/cri
         by:



WARN: grpc-ecosystem/go-grpc-prometheus has 3 versions imported
        v1.2.0 by:
                docker/docker
        v1.1 by:
                containerd/cri
        6b7015e65d366bf3f19b2b2a000a831940f0f7e0 by:
                docker/swarmkit
                containerd/containerd


WARN: opencontainers/image-spec has 2 versions imported
        ab7389ef9f50030c9b245bc16b981c7ddf192882 by:
                docker/distribution
        v1.0.1 by:
                docker/docker
                containerd/cri
                moby/buildkit
                docker/libnetwork
                docker/swarmkit
                containerd/containerd


WARN: gogo/googleapis has 2 versions imported
        08a7655d27152912db7aaf4f983275eaf8d128ef by:
                docker/docker
                containerd/cri
                containerd/containerd
        b23578765ee54ff6bceff57f397d833bf4ca6869 by:
                moby/buildkit


WARN: syndtr/gocapability has 2 versions imported
        2c00daeb6c3b45114c80ac44119e7b8801fdd852 by:
                docker/docker
        db04d3cc01c8b54962a58ec7e491717d06cfcc16 by:
                containerd/cri
                opencontainers/runc
                moby/buildkit
                containerd/containerd


WARN: spf13/pflag has 4 versions imported
        v1.0.1 by:
                docker/docker
        4c012f6dcd9546820e378d0bdda4d8fc772cdfea by:
                containerd/continuity
        5644820622454e71517561946e3d94b9f9db6842 by:
                docker/distribution
        7f60f83a2c81bc3c3c0d5297f61ddfa68da9d3b7 by:
                docker/swarmkit


WARN: urfave/cli has 2 versions imported
        d53eb991652b1d438abdd34ce4bfa3ef1539108e by:
                opencontainers/runc
        7bc6a0acffa589f415f88aca16cc1de5ffd66f9c by:
                containerd/cri
                moby/buildkit
                containerd/containerd


WARN: docker/distribution has 3 versions imported
        83389a148052d74ac602f5f1d62f86ff2f3c4aa5 by:
                docker/docker
                docker/swarmkit
        b38e5838b7b2f2ad48e06ec4b500011976080621 by:
                containerd/cri
                containerd/containerd
        30578ca32960a4d368bf6db67b0a33c2a1f3dc6f by:
                moby/buildkit


WARN: pborman/uuid has 2 versions imported
        v1.0 by:
                docker/docker
        c65b2f87fee37d1c7854c9164a450713c28d50cd by:
                containerd/containerd


WARN: matttproud/golang_protobuf_extensions has 2 versions imported
        v1.0.0 by:
                docker/docker
                containerd/cri
                containerd/containerd
                docker/swarmkit
        c12348ce28de40eed0136aa2b644d0ee0650e56c by:
                docker/distribution


WARN: docker/go-metrics has 3 versions imported
        399ea8c73916000c64c2c76e8da00ca82f8387ab by:
                docker/distribution
        d466d4f6fd960e01820085bd7e1a24426ee7ef18 by:
                docker/docker
                docker/swarmkit
        4ea375f7759c82740c893fc030bc37088d2ec098 by:
                containerd/cri
                containerd/containerd


WARN: docker/libkv has 2 versions imported
        458977154600b9f23984d9f4b82e79570b5ae12b by:
                docker/docker
                docker/libnetwork
        1d8431073ae03cdaedb198a89722f3aab6d418ef by:
                docker/swarmkit


WARN: coreos/etcd has 2 versions imported
        v3.3.9 by:
                docker/docker
        v3.2.1 by:
                docker/libnetwork
                docker/swarmkit


WARN: prometheus/common has 3 versions imported
        7600349dcfe1abd18d72d3a1770870d9800a7801 by:
                docker/docker
        89604d197083d4781071d3c65855d24ecfb0a563 by:
                containerd/cri
                docker/distribution
                containerd/containerd
        ebdfc6da46522d58825777cf1f90490a5b1ef1d8 by:
                docker/swarmkit


Summary:
        - 160 repo(s) under github.com were scanned in 2.0020956s.
        - 58 warning(s) were found.
        - 58 repo(s) are imported at different revisions.
        - 108 repo(s) were skipped
```