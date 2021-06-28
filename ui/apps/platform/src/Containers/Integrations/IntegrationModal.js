import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import { actions } from 'reducers/integrations';
import { createStructuredSelector } from 'reselect';
import * as Icon from 'react-feather';
import { selectors } from 'reducers';

import Dialog from 'Components/Dialog';
import Form from 'Containers/Integrations/Form';
import Modal from 'Components/Modal';
import { toggleRow, toggleSelectAll } from 'utils/checkboxUtils';
import { getDefaultValues, setStoredCredentialFields } from './utils/integrationFormUtils';

const SOURCE_LABELS = Object.freeze({
    authPlugins: 'auth plugins',
    authProviders: 'authentication provider',
    imageIntegrations: 'image integrations',
    notifiers: 'notifier integrations',
    backups: 'backup integrations',
});

class IntegrationModal extends Component {
    static propTypes = {
        integrations: PropTypes.arrayOf(
            PropTypes.shape({
                type: PropTypes.string,
            })
        ).isRequired,
        source: PropTypes.oneOf([
            'authPlugins',
            'imageIntegrations',
            'notifiers',
            'authProviders',
            'backups',
        ]).isRequired,
        type: PropTypes.string.isRequired,
        label: PropTypes.string.isRequired,
        onRequestClose: PropTypes.func.isRequired,
        deleteIntegrations: PropTypes.func.isRequired,
        isCreating: PropTypes.bool,
        setCreateState: PropTypes.func.isRequired,
        selectedIntegration: PropTypes.shape({}).isRequired,
    };

    static defaultProps = {
        isCreating: false,
    };

    constructor(props) {
        super(props);

        this.state = {
            showAutogenerated: false,
            showConfirmationDialog: false,
            onlyOneIntegrationAllowed: false,
            selection: [],
        };
    }

    componentWillUnmount() {
        this.props.setCreateState(false);
    }

    onTableDelete = ({ id }) => {
        const { length } = this.state.selection;
        const { source, type } = this.props;
        if (length) {
            this.showConfirmationDialog();
        } else {
            this.props.deleteIntegrations(source, type, [id]);
        }
    };

    onTableAdd = (options) => {
        if (options?.onlyOneIntegrationAllowed) {
            this.setState({ showConfirmationDialog: true, onlyOneIntegrationAllowed: true });
        } else {
            this.props.setCreateState(true);
        }
    };

    onAutogeneratedClick = () => {
        this.setState((prevState) => ({ showAutogenerated: !prevState.showAutogenerated }));
    };

    onTableRowClick = (integration) => {
        this.setState({
            selectedIntegration: integration,
        });
        this.props.setCreateState(false);
    };

    setTableRef = (table) => {
        this.integrationTable = table;
    };

    getSelectedIntegrationId = () =>
        this.state.selectedIntegration ? this.state.selectedIntegration.id : '';

    // determines whether the form panel is open based on selected integration and creation state
    formIsOpen = () => this.props.isCreating || !!this.state.selectedIntegration;

    hideConfirmationDialog = () => {
        this.setState({ showConfirmationDialog: false, onlyOneIntegrationAllowed: false });
    };

    showConfirmationDialog = () => {
        this.setState({ showConfirmationDialog: true });
    };

    closeIntegrationForm = () => {
        this.props.onRequestClose();
    };

    clearSelection = () => this.setState({ selection: [] });

    activateAuthIntegration = (integration) => {
        if (integration !== null && integration.loginUrl !== null && !integration.validated) {
            window.location = integration.loginUrl;
        }
    };

    deleteTableSelectedIntegrations = () => {
        const { selection } = this.state;
        const { source, type } = this.props;
        this.props.deleteIntegrations(source, type, selection);
        this.clearSelection();
        this.hideConfirmationDialog();
    };

    toggleRow = (id) => {
        const selection = toggleRow(id, this.state.selection);
        this.updateSelection(selection);
    };

    toggleSelectAll = () => {
        const rowsLength = this.props.integrations.length;
        const tableRef = this.integrationTable.reactTable;
        const selection = toggleSelectAll(rowsLength, this.state.selection, tableRef);
        this.updateSelection(selection);
    };

    updateSelection = (selection) => this.setState({ selection });

    renderHeader = () => {
        const { source, type } = this.props;
        let { label } = this.props;
        if (label === undefined) {
            label = type;
        }
        return (
            <header className="flex items-center w-full p-4 bg-primary-600 text-base-100 uppercase">
                <span className="flex flex-1">{`Configure ${label} ${SOURCE_LABELS[source]}`}</span>
                <Icon.X className="h-4 w-4 cursor-pointer" onClick={this.props.onRequestClose} />
            </header>
        );
    };

    renderForm = () => {
        const { source, type, selectedIntegration } = this.props;

        const initialValues = selectedIntegration || getDefaultValues(source, type);
        // we want to add a new "hasStoredCredentials" field to determine whether this
        // integration can possibly use stored credentials if the password field is empty
        // on submission
        const modifiedSelectedIntegration = setStoredCredentialFields(source, type, initialValues);
        return <Form initialValues={modifiedSelectedIntegration} source={source} type={type} />;
    };

    renderConfirmationDialog = () => {
        // TODO: replace form descriptors with something with less indirection
        if (this.state.onlyOneIntegrationAllowed) {
            return (
                <Dialog
                    isOpen={this.state.showConfirmationDialog}
                    text="Only one integration of this type is allowed. Edit the existing integration, or delete it and then create a new one."
                    onCancel={this.hideConfirmationDialog}
                    cancelText="OK"
                />
            );
        }

        const numSelectedRows = this.state.selection.length;
        return (
            <Dialog
                isOpen={this.state.showConfirmationDialog}
                text={`Are you sure you want to delete ${numSelectedRows} integration(s)?`}
                onConfirm={this.deleteTableSelectedIntegrations}
                onCancel={this.hideConfirmationDialog}
            />
        );
    };

    render() {
        const headerSection = this.renderHeader();
        const formSection = this.renderForm();
        const confirmationDialog = this.renderConfirmationDialog();

        return (
            <Modal isOpen onRequestClose={this.props.onRequestClose} className="w-1/2">
                {headerSection}
                <div className="flex h-full relative w-full">{formSection}</div>
                {confirmationDialog}
            </Modal>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    isCreating: selectors.getCreationState,
});

const mapDispatchToProps = (dispatch) => ({
    deleteIntegrations: (source, sourceType, ids) =>
        dispatch(actions.deleteIntegrations(source, sourceType, ids)),
    setCreateState: (state) => dispatch(actions.setCreateState(state)),
});

export default connect(mapStateToProps, mapDispatchToProps)(IntegrationModal);
