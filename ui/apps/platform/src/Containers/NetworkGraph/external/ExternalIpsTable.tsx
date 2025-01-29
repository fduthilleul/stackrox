import React, { useCallback } from 'react';
import { Flex, Pagination, Toolbar, ToolbarContent, ToolbarItem } from '@patternfly/react-core';
import { InnerScrollContainer, Table, Tbody, Td, Th, Thead, Tr } from '@patternfly/react-table';

import ExternalLink from 'Components/PatternFly/IconText/ExternalLink';
import TbodyUnified from 'Components/TableStateTemplates/TbodyUnified';
import useMetadata from 'hooks/useMetadata';
import useRestQuery from 'hooks/useRestQuery';
import useURLPagination from 'hooks/useURLPagination';
import { getExternalIpsFlowsMetadata } from 'services/NetworkService';
import { getTableUIState } from 'utils/getTableUIState';
import { getVersionedDocs } from 'utils/versioning';
import { ExternalNetworkFlowsMetadataResponse } from 'types/networkFlow.proto';

import { NetworkScopeHierarchy } from '../types/networkScopeHierarchy';

export type ExternalIpsTableProps = {
    scopeHierarchy: NetworkScopeHierarchy;
};

function ExternalIpsTable({ scopeHierarchy }: ExternalIpsTableProps) {
    const { version } = useMetadata();
    const pagination = useURLPagination(10);
    const { page, perPage, setPage, setPerPage } = pagination;
    const clusterId = scopeHierarchy.cluster.id;
    const { namespaces, deployments } = scopeHierarchy;
    const fetchExternalIpsFlowsMetadata =
        useCallback((): Promise<ExternalNetworkFlowsMetadataResponse> => {
            return getExternalIpsFlowsMetadata(clusterId, namespaces, deployments, {
                sortOption: {},
                page,
                perPage,
                advancedFilters: {},
            });
        }, [page, perPage, clusterId, deployments, namespaces]);

    const {
        data: externalIpsFlowsMetadata,
        isLoading,
        error,
    } = useRestQuery(fetchExternalIpsFlowsMetadata);

    const tableState = getTableUIState({
        isLoading,
        data: externalIpsFlowsMetadata?.entities,
        error,
        searchFilter: {},
    });

    return (
        <>
            <Toolbar>
                <ToolbarContent>
                    <ToolbarItem variant="pagination" align={{ default: 'alignRight' }}>
                        <Pagination
                            itemCount={externalIpsFlowsMetadata?.totalEntities ?? 0}
                            page={page}
                            perPage={perPage}
                            onSetPage={(_, newPage) => setPage(newPage)}
                            onPerPageSelect={(_, newPerPage) => setPerPage(newPerPage)}
                            isCompact
                        />
                    </ToolbarItem>
                </ToolbarContent>
            </Toolbar>
            <InnerScrollContainer>
                <Table variant="compact">
                    <Thead>
                        <Tr>
                            <Th>Entity</Th>
                            <Th>Internal flows</Th>
                        </Tr>
                    </Thead>
                    <TbodyUnified
                        tableState={tableState}
                        colSpan={7}
                        errorProps={{
                            title: 'There was an error loading external ips',
                        }}
                        emptyProps={{
                            message: 'No external ips found. This feature might not be enabled.',
                            children: (
                                <Flex alignSelf={{ default: 'alignSelfCenter' }}>
                                    <ExternalLink>
                                        <a
                                            href={getVersionedDocs(
                                                version,
                                                'operating/visualizing-external-entities'
                                            )}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                        >
                                            Enabling external ip collection
                                        </a>
                                    </ExternalLink>
                                </Flex>
                            ),
                        }}
                        renderer={({ data }) => (
                            <Tbody>
                                {data.map(({ entity, flowsCount }) => {
                                    return (
                                        <Tr key={entity.id}>
                                            <Td dataLabel="Entity">{entity.externalSource.name}</Td>
                                            <Td dataLabel="Internal flows">{flowsCount}</Td>
                                        </Tr>
                                    );
                                })}
                            </Tbody>
                        )}
                    />
                </Table>
            </InnerScrollContainer>
        </>
    );
}

export default ExternalIpsTable;
