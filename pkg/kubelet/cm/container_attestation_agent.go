import (
	"fmt"
	"sync"
	"time"
)

type nsMeasurements {
	nsMap map[string]string
	mutex sync.Mutex
}

nsMappings := nsMeasurments{ nsMap : make(map[string]string)}

/*
 * Add NS to dictionary
 * Associate container with NS
 */
func attestNewNS(ns string) {

	nsMappings.mutex.Lock()

	nsMappings.map[ns] := ""

	// Check the log and gather measurements 
	// from before we got the NS (likely containerd runc 
	// measurements??)

	nsMappings.mutex.Unlock()

	return 

}

/*
 * Attestation: parseConatainerLinuxNS
 *      Get Linux NS of container process to match 
 *      OS-level perspective for IMA log parsing
 */
func parseContainerLinuxNS(pid int) string, err {

        cgroupPath := ""
        fmt.Sprintf(cgroupPath,"/proc/%d/ns/cgroup", pid)

        symlink, err := os.Readlink(cgroupPath)
        if err != nil {
                return err
        }

        split := strings.Split(symlink, "[")

        split = strings.Split(split[1], "]")

        ns := split[0]


        return ns
}
/*
 * Attestation: registerAttestationAgent
 *      Get Linux NS of new container process
 *      and tell Keylime this information 
 *      to parse IMA logs for container
 */
func registerAttestationAgent(pid int) {

        ns, err := parseContainerLinuxNS(pid)
        if err != nil {
                return
        }


        // Tell keylime NS

        return attestNewNS(ns)


}

