/*
Copyright 2019 The Tekton Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import React from 'react';
import { waitForElement, fireEvent, cleanup } from 'react-testing-library';
import { renderWithRouter } from '../../../test/utils/test';
import 'react-testing-library/cleanup-after-each';
import { WebhookCreate } from '../WebhookCreate';
import * as API from '../../../api/index';
import 'jest-dom/extend-expect'

global.scrollTo = jest.fn();

const namespaces = ["default", "istio-system", "namespace3"];

const pipelines = [
  {
    metadata: {
      name: "pipeline0",
      namespace: "default",
    },
  },
  {
    metadata: {
      name: "simple-pipeline",
      namespace: "default",
    },
  },
  {
    metadata: {
      name: "simple-helm-pipeline-insecure",
      namespace: "istio-system",
    }
  }
];

const secretsResponseMock = [
  {
    "name": "ghe",
  },
  {
    "name": "git",
  }
]

const secretsDeletedMock = [
  {
    "name": "git",
  }
]

const deleteSecretSuccessMock = {}

const serviceAccounts = [
  {
    metadata: {
      name: "default",
      namespace: "default"
    }
  },
  {
    metadata: {
      name: "testserviceaccount",
      namespace: "istio-system",
    },
  }
];

beforeEach(() => {
  jest.restoreAllMocks
  jest.resetModules()
 });

afterEach(() => {
  jest.clearAllMocks()
  cleanup()
 });

//-----------------------------------//
describe('confirm deletion success', () => {
  it('delete button should hide modal, remove secret from listing and reset dropdown', async () => {
    jest.spyOn(API, 'getSecrets').mockImplementation(() => Promise.resolve(secretsResponseMock));
    jest.spyOn(API, 'deleteSecret').mockImplementation(() => Promise.resolve(deleteSecretSuccessMock));
    const { getByText } = renderWithRouter(
      <WebhookCreate
        match={{}}
        namespaces={namespaces}
        pipelines={pipelines}
        fetchPipelines={() => {}}
        serviceAccounts={serviceAccounts}
        fetchServiceAccounts={() => {}}
        setShowNotificationOnTable={() => {}}
      />
    );
    fireEvent.click(await waitForElement(() => getByText(/select namespace/i)));
    fireEvent.click(await waitForElement(() => getByText(/istio-system/i)));
    await waitForElement(() => document.getElementsByClassName('secButtonEnabled'));
    fireEvent.click(await waitForElement(() => getByText(/select secret/i)));
    fireEvent.click(await waitForElement(() => getByText(/ghe/i)));
    await waitForElement(() => document.getElementsByClassName('secButtonEnabled'));
    fireEvent.click(document.getElementById('delete-secret-button'));
    expect(document.getElementById('delete-modal').getAttribute('class')).toContain('is-visible');

    jest.spyOn(API, 'getSecrets').mockImplementation(() => Promise.resolve(secretsDeletedMock));
    fireEvent.click(document.getElementById('delete-modal').getElementsByClassName('bx--btn--danger').item(0));

    await waitForElement(() => getByText(/Secret deleted./));
    expect(document.getElementById('delete-modal').getAttribute('class')).not.toContain('is-visible');
    await waitForElement(() => getByText(/select secret/i));
  });

})
