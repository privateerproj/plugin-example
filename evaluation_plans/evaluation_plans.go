package evaluation_plans

import (
    "github.com/ossf/gemara/layer4"
    "github.com/MYORG/plugin-GCS/evaluation_plans/reusable_steps"
)

var (
    // Add more entries here if other catalogs or catalog versions are adopted by the plugin
    // Then use orchestrator.AddEvaluationSuite to make these available to the user
    CCC_ObjStor = map[string][]layer4.AssessmentStep{
        // For best results, ensure every assessment id is represented in this map
        "CCC.ObjStor.CN01.AR01": {
            /* When a request is made to read a bucket, the service
MUST prevent any request using KMS keys not listed as trusted by
the organization.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN01.AR02": {
            /* When a request is made to read an object, the service
MUST prevent any request using KMS keys not listed as trusted by
the organization.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN01.AR03": {
            /* When a request is made to write to a bucket, the service MUST
prevent any request using KMS keys not listed as trusted by the
organization.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN01.AR04": {
            /* When a request is made to write to an object, the service MUST
prevent any request using KMS keys not listed as trusted by the
organization.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN03.AR01": {
            /* When an object storage bucket deletion is attempted, the bucket MUST be
fully recoverable for a set time-frame after deletion is requested.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN03.AR02": {
            /* When an attempt is made to modify the retention policy for an object
storage bucket, the service MUST prevent the policy from being modified.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN04.AR01": {
            /* When an object is uploaded to the object storage system, the object
MUST automatically receive a default retention policy that prevents
premature deletion or modification.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN04.AR02": {
            /* When an attempt is made to delete or modify an object that is subject
to an active retention policy, the service MUST prevent the action
from being completed.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN05.AR01": {
            /* When an object is uploaded to the object storage bucket, the object
MUST be stored with a unique identifier.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN05.AR02": {
            /* When an object is modified, the service MUST assign a new unique
identifier to the modified object to differentiate it from the
previous version.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN05.AR03": {
            /* When an object is modified, the service MUST allow for recovery
of previous versions of the object.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN05.AR04": {
            /* When an object is deleted, the service MUST retain other versions of
the object to allow for recovery of previous versions.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN06.AR01": {
            /* When an object storage bucket is accessed, the service MUST store
access logs in a separate data store.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN06.AR01": {
            /* When an object storage bucket stores access logs for other object
storage buckets, the bucket MUST be classified as the highest
possible sensitivity level.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN01.AR01": {
            /* When a port is exposed for non-SSH network traffic, all traffic
MUST include a TLS handshake AND be encrypted using TLS 1.3 or
higher.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN01.AR02": {
            /* When a port is exposed for SSH network traffic, all traffic MUST
include a SSH handshake AND be encrypted using SSHv2 or higher.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN01.AR03": {
            /* When the service receives unencrypted traffic, 
then it MUST either block the request or automatically
redirect it to the secure equivalent.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN01.AR07": {
            /* When a port is exposed, the service MUST ensure that the protocol
and service officially assigned to that port number by the IANA
Service Name and Transport Protocol Port Number Registry, and no
other, is run on that port.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN01.AR08": {
            /* When a service transmits data using TLS, mutual TLS (mTLS) MUST be
implemented to require both client and server certificate
authentication for all connections.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN02.AR01": {
            /* When data is stored, it MUST be encrypted using the latest
industry-standard encryption methods.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN06.AR01": {
            /* When the service is running, its region and availability zone MUST
be included in a list of explicitly trusted or approved locations
within the trust perimeter.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN06.AR02": {
            /* When a child resource is deployed, its region and availability
zone MUST be included in a list of explicitly trusted or approved
locations within the trust perimeter.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN09.AR01": {
            /* When the service is operational, its logs and any child resource
logs MUST NOT be accessible from the resource they record access
to.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN09.AR02": {
            /* When the service is operational, disabling the logs for the service
or its child resources MUST NOT be possible without also disabling
the corresponding resource.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN09.AR03": {
            /* When the service is operational, any attempt to redirect logs for
the service or its child resources MUST NOT be possible without
halting operation of the corresponding resource and publishing
corresponding events to monitored channels.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN10.AR01": {
            /* When data is replicated, the service MUST ensure that replication
only occurs to destinations that are explicitly included within
the defined trust perimeter.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN02.AR01": {
            /* When a permission set is allowed for an object in a bucket, the
service MUST allow the same permission set to access all objects
in the same bucket.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.ObjStor.CN02.AR02": {
            /* When a permission set is denied for an object in a bucket, the
service MUST deny the same permission set to access all objects
in the same bucket.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN03.AR01": {
            /* When an entity attempts to modify the service through a user
interface, the authentication process MUST require multiple
identifying factors for authentication.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN03.AR02": {
            /* When an entity attempts to modify the service through an API
endpoint, the authentication process MUST require a credential
such as an API key or token AND originate from within the trust
perimeter.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN03.AR03": {
            /* When an entity attempts to view information on the service through
a user interface, the authentication process MUST require multiple
identifying factors from the user.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN03.AR04": {
            /* When an entity attempts to view information on the service through
an API endpoint, the authentication process MUST require a
credential such as an API key or token AND originate from within
the trust perimeter.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR01": {
            /* When an attempt is made to modify data on the service or a child
resource, the service MUST block requests from unauthorized
entities.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR02": {
            /* When administrative access or configuration change is attempted on
the service or a child resource, the service MUST refuse requests
from unauthorized entities.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR03": {
            /* When administrative access or configuration change is attempted on
the service or a child resource in a multi-tenant environment, the
service MUST refuse requests across tenant boundaries unless the
origin is explicitly included in a pre-approved allowlist.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR04": {
            /* When data is requested from outside the trust perimeter, the
service MUST refuse requests from unauthorized entities.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR05": {
            /* When any request is made from outside the trust perimeter,
the service MUST NOT provide any response that may indicate the
service exists.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05.AR06": {
            /* When any request is made to the service or a child resource, the
service MUST refuse requests from unauthorized entities.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN04.AR01": {
            /* When administrative access or configuration change is attempted on
the service or a child resource, the service MUST log the client
identity, time, and result of the attempt.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN04.AR02": {
            /* When any attempt is made to modify data on the service or a child
resource, the service MUST log the client identity, time, and
result of the attempt.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN04.AR03": {
            /* When any attempt is made to read data on the service or a child
resource, the service MUST log the client identity, time, and
result of the attempt.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN07.AR01": {
            /* When enumeration activities are detected, the service MUST publish
an event to a monitored channel which includes the client
identity, time, and nature of the activity.
 */
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN07.AR02": {
            /* When enumeration activities are detected, the service MUST log the
client identity, time, and nature of the activity.
 */
            reusable_steps.NotImplemented,
        },
        
        // CCC
        "CCC.Core.CN01": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN02": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN03": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN04": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN05": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN06": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN07": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN09": {
            reusable_steps.NotImplemented,
        },
        "CCC.Core.CN10": {
            reusable_steps.NotImplemented,
        },
    }
)
