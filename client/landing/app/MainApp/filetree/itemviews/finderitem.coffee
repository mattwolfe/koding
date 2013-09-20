class NFinderItem extends JTreeItemView

  constructor:(options = {},data)->

    options.tagName or= "li"
    options.type    or= "finderitem"

    super options, data

    @isLoading        = no
    @beingDeleted     = no
    @beingEdited      = no

    childConstructor = switch data.type
      when "vm"         then NVMItemView
      when "folder"     then NFolderItemView
      when "section"    then NSectionItemView
      when "mount"      then NMountItemView
      when "brokenLink" then NBrokenLinkItemView
      else NFileItemView

    @childView = new childConstructor delegate: this, data
    @childView.$().css "margin-left", (data.depth)*10

    if data.name? and data.name.length > 20 - data.depth
      @childView.setDomAttributes
        title : FSHelper.plainPath data.name

    @progress = new KDProgressBarView

    @on "ItemBeingDeleted", =>
      data.removeLocalFileInfo()

    @on "viewAppended", =>
      fileInfo = data.getLocalFileInfo()
      if fileInfo.lastUploadedChunk
        {lastUploadedChunk, totalChunks} = fileInfo
        if lastUploadedChunk is totalChunks
          data.removeLocalFileInfo()
          @updateProgressView 100
        else
          @updateProgressView 100 * lastUploadedChunk / totalChunks

  mouseDown:-> yes

  resetView:(view)->

    if @deleteView
      @deleteView.destroy()
      delete @deleteView

    if @renameView
      @renameView.destroy()
      delete @renameView

    @childView.show()
    @beingDeleted = no
    @beingEdited = no
    @callback = null
    @unsetClass "being-deleted being-edited"
    @getDelegate().setKeyView()

  confirmDelete:(callback)->

    @callback = callback
    @showDeleteView()

  showDeleteView:->

    return if @deleteView
    @setClass "being-deleted"
    @beingDeleted = yes
    @childView.hide()
    data = @getData()
    @addSubView @deleteView = new NFinderItemDeleteView {}, data
    @deleteView.on "FinderDeleteConfirmation", (confirmation)=>
      @callback? confirmation
      @resetView()
    @deleteView.setKeyView()

  showRenameView:(callback)->

    return if @renameView
    @setClass "being-edited"
    @beingEdited = yes
    @callback = callback
    @childView.hide()
    data = @getData()
    @addSubView @renameView = new NFinderItemRenameView {}, data
    @renameView.$().css "margin-left", ((data.depth+1)*10)+2
    @renameView.on "FinderRenameConfirmation", (newValue)=>
      @callback? newValue
      @resetView()
    @renameView.input.setFocus()

  updateProgressView: (percent=0, determinate=yes)->
    @progress.setOption "determinate", determinate
    @progress.updateBar percent, "%", ""
    if 0 <= percent < 100
    then @setClass   "progress"
    else @utils.wait 1000, => @unsetClass "progress"

  pistachio:->

    """
    {{> @childView}}
    {{> @progress}}
    """
