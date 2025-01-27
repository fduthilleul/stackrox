import React, { useState } from 'react';
import { Divider, Stack, StackItem, ToggleGroup, ToggleGroupItem } from '@patternfly/react-core';
import useFeatureFlags from 'hooks/useFeatureFlags';
import { CustomNodeModel } from '../types/topology.type';
import { EdgeState } from '../components/EdgeStateSelect';
import { Flow } from '../types/flow.type';
import InternalFlows from './InternalFlows';
import ExternalFlows from './ExternalFlows';

export type DeploymentFlowsView = 'external-flows' | 'internal-flows';

type DeploymentFlowsProps = {
    deploymentId: string;
    nodes: CustomNodeModel[];
    edgeState: EdgeState;
    onNodeSelect: (id: string) => void;
    isLoadingNetworkFlows: boolean;
    networkFlowsError: string;
    networkFlows: Flow[];
    refetchFlows: () => void;
};

function DeploymentFlows({
    deploymentId,
    nodes,
    edgeState,
    onNodeSelect,
    isLoadingNetworkFlows,
    networkFlowsError,
    networkFlows,
    refetchFlows,
}: DeploymentFlowsProps) {
    const { isFeatureFlagEnabled } = useFeatureFlags();
    const isExternalIpsEnabled = isFeatureFlagEnabled('ROX_EXTERNAL_IPS');
    const [selectedView, setSelectedView] = useState<DeploymentFlowsView>('internal-flows');

    if (!isExternalIpsEnabled) {
        return (
            <div className="pf-v5-u-h-100 pf-v5-u-p-md">
                <InternalFlows
                    nodes={nodes}
                    deploymentId={deploymentId}
                    edgeState={edgeState}
                    onNodeSelect={onNodeSelect}
                    isLoadingNetworkFlows={isLoadingNetworkFlows}
                    networkFlowsError={networkFlowsError}
                    networkFlows={networkFlows}
                    refetchFlows={refetchFlows}
                />
            </div>
        );
    }

    return (
        <div className="pf-v5-u-h-100">
            <Stack>
                <StackItem className="pf-v5-u-p-md">
                    <ToggleGroup aria-label="Toggle between internal flows and external flows views">
                        <ToggleGroupItem
                            text="Internal flows"
                            buttonId="internal-flows"
                            isSelected={selectedView === 'internal-flows'}
                            onChange={() => setSelectedView('internal-flows')}
                        />
                        <ToggleGroupItem
                            text="External flows"
                            buttonId="external-flows"
                            isSelected={selectedView === 'external-flows'}
                            onChange={() => setSelectedView('external-flows')}
                        />
                    </ToggleGroup>
                </StackItem>
                <Divider component="hr" />
                <StackItem isFilled style={{ overflow: 'auto' }}>
                    <Stack className="pf-v5-u-p-md">
                        {selectedView === 'internal-flows' ? (
                            <InternalFlows
                                nodes={nodes}
                                deploymentId={deploymentId}
                                edgeState={edgeState}
                                onNodeSelect={onNodeSelect}
                                isLoadingNetworkFlows={isLoadingNetworkFlows}
                                networkFlowsError={networkFlowsError}
                                networkFlows={networkFlows}
                                refetchFlows={refetchFlows}
                            />
                        ) : (
                            <ExternalFlows deploymentId={deploymentId} />
                        )}
                    </Stack>
                </StackItem>
            </Stack>
        </div>
    );
}

export default DeploymentFlows;
