import { gql } from '@apollo/client';

import { VulnerabilitySeverity } from 'types/cve.proto';

// This type is specific to the way we query using GraphQL
export type Vulnerability = {
    id: string;
    cve: string;
    isFixable: boolean;
    severity: VulnerabilitySeverity;
    cvss: string;
    scoreVersion: string;
    discoveredAtImage: string;
    components: EmbeddedImageScanComponent[];
};

// This type is specific to the way we query using GraphQL
export type EmbeddedImageScanComponent = { id: string; name: string; fixedIn: string };

export type GetObservedCVEsData = {
    result: {
        name: {
            fullName: string;
            tag: string;
        };
        vulnCount: number;
        vulns: Vulnerability[];
    };
};

export type GetObservedCVEsVars = {
    imageId: string;
    vulnsQuery: string;
};

export const GET_OBSERVED_CVES = gql`
    query getObservedCVEs($imageId: ID!, $vulnsQuery: String) {
        result: image(id: $imageId) {
            name {
                fullName
                tag
            }
            vulnCount(query: $vulnsQuery)
            vulns(query: $vulnsQuery) {
                id: cve
                cve
                isFixable
                severity
                scoreVersion
                cvss
                discoveredAtImage
                components {
                    id
                    name
                    fixedIn
                }
            }
        }
    }
`;
