import os
from isort import file
import yaml
import mdutils
from mdutils.mdutils import MdUtils

CRDS_DIR = "package/crds"
DOCS_DIR = "docs/"
TPL_DIR = "docs/templates/"


def open_yaml(filename):
    with open(file=filename, mode='r') as stream:
        try:
            data = yaml.safe_load(stream)
            return data
        except yaml.YAMLError as exc:
            print(exc)
            exit(1)


def readfile(filename):
    with open(filename, 'r') as file:
        return file.read()


def contains(l, elt):
    for e in l:
        if e == elt:
            return True
    return False


def is_field_present(block, fieldName):
    try:
        _ = block[fieldName]
        return True
    except KeyError:
        return False


def is_optional(properties, required):
    return len(properties) != len(required)


def recursive_read(mdFile, block, field, count=0):
    if is_field_present(block[field], 'properties'):
        required = is_field_present(block[field], 'required')
        if required:
            for f in block[field]['properties']:
                if contains(block[field]['required'], f):
                    mdFile.write("\n"+count*2*" "+"* `"+f+"`: " +
                                 block[field]['properties'][f]['description'])
                    mdFile.write(" (Required)", color='orange')

                    recursive_read(
                        mdFile, block[field]['properties'], f, count+1)

                    mdFile.new_line()

        for f in block[field]['properties']:
            if not required or not contains(block[field]['required'], f):
                mdFile.write("\n"+count*2*" "+"* `"+f+"`: " +
                             block[field]['properties'][f]['description'])

                recursive_read(mdFile, block[field]['properties'], f, count+1)

                mdFile.new_line()
    else:
        return


def createMdDocs(data, filename):
    mdFile = MdUtils(file_name=DOCS_DIR+filename, title=filename)

    mdFile.new_paragraph("This document has been generated from the CRD.\n")
    mdFile.new_line()

    mdFile.new_header(1, "Example")
    mdFile.insert_code(readfile(TPL_DIR+filename), language="yaml")
    mdFile.new_line()

    mdFile.new_header(1, "Schema")
    mdFile.new_line()

    required = True
    try:
        if len(data['required']) != 0:
            mdFile.new_header(2, "Required")
            mdFile.new_line()

            for field in data['properties']:
                if contains(data['required'], field):
                    mdFile.write("`"+field+"`: " +
                                 data['properties'][field]['description']+"\n")

                    recursive_read(mdFile, data['properties'], field)

                    mdFile.new_line()
    except KeyError:
        required = False

    if required and not is_optional(data['properties'], data['required']):
        mdFile.create_md_file()
        return

    mdFile.new_header(2, "Optional")
    mdFile.new_line()

    for field in data['properties']:
        if not required or not contains(data['required'], field):
            mdFile.write("`"+field+"`: " +
                         data['properties'][field]['description']+"\n")

            recursive_read(mdFile, data['properties'], field)

            mdFile.new_line()

    mdFile.create_md_file()


def generate_resource_docs(filename):
    print("generating docs for: ", filename)
    data = open_yaml(filename)

    resource_data = data['spec']['versions'][0]['schema']['openAPIV3Schema'][
        'properties']['spec']['properties']['forProvider']

    mdfilename = data['spec']['group']
    createMdDocs(resource_data, mdfilename)


def generate_provider_docs(filename):
    print("generating docs for: ", filename)
    data = open_yaml(filename)
    if f == "package/crds/taikun.jet.crossplane.io_providerconfigusages.yaml":
        return
        provider_data = data['spec']['versions'][0]['schema']['openAPIV3Schema']
        mdfilename = data['metadata']['name']
        createMdDocs(provider_data, mdfilename)
    elif f == "package/crds/taikun.jet.crossplane.io_providerconfigs.yaml":
        provider_data = data['spec']['versions'][0]['schema']['openAPIV3Schema']['properties']['spec']
        mdfilename = data['metadata']['name']
        createMdDocs(provider_data, mdfilename)
    elif f == "package/crds/taikun.jet.crossplane.io_storeconfigs.yaml":
        return
        provider_data = data['spec']['versions'][0]['schema']['openAPIV3Schema']['properties']['spec']
        mdfilename = data['metadata']['name']
        createMdDocs(provider_data, mdfilename)
    else:
        return

#
# --------------------------------------------------------------------------------------------------------------

# --------------------------------------------------------------------------------------------------------------
#


if not os.path.exists(path=CRDS_DIR):
    print("CRD's do not exists.")


for filename in os.listdir(CRDS_DIR):
    f = os.path.join(CRDS_DIR, filename)
    if os.path.isfile(f) and (filename.endswith('.yaml') or filename.endswith('.yml')) and not filename.startswith('taikun.'):
        generate_resource_docs(f)
    elif os.path.isfile(f) and (filename.endswith('.yaml') or filename.endswith('.yml')) and filename.startswith('taikun.'):
        generate_provider_docs(f)
    else:
        continue
